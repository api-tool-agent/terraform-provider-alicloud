package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/stretchr/testify/assert"
)

// newTestServerError builds a synthetic SDK ServerError with the given HTTP status
// and Alibaba Cloud error code, mirroring how the real SDK deserializes an error
// response (the error code is parsed out of the JSON body's "Code" field). It
// keeps message text free of any code substring so IsExpectedErrors' message
// matching cannot produce false positives in the assertions below.
func newTestServerError(t *testing.T, httpStatus int, code, message string) *errors.ServerError {
	t.Helper()
	body := fmt.Sprintf(`{"Code":%q,"Message":%q}`, code, message)
	err := errors.NewServerError(httpStatus, body, "")
	se, ok := err.(*errors.ServerError)
	if !ok {
		t.Fatalf("NewServerError did not return *ServerError: %T", err)
	}
	return se
}

// TestRDSDBInstanceGoneStatusClassification exercises the error-classification
// logic that the recycle-bin / refunded-instance fix relies on.
//
// When a Subscription (pay-by-month/year) RDS instance is unsubscribed out of band via the
// console, it enters refund -> lock -> release -> recycle-bin. During the
// recycle-bin retention window the instance object still exists and the RDS API
// returns 403 OperationDenied.DBInstanceStatus / OperationDenied.ReadDBInstanceStatus
// — NOT a 404. The provider previously treated those 403s as retryable-then-
// hard-fail, so `terraform refresh` died and `terraform destroy` hung. The fix
// elevates these two codes to a first-class "parent instance gone" signal that is
// equal to NotFound: DescribeDBInstance maps them to NotFound, and child-resource
// Read/Delete paths confirm via a follow-up DescribeDBInstance before dropping state.
//
// This test pins the classification primitives the fix depends on:
//   - A terminal 403 (either gone code) is matched by
//     IsExpectedErrors(err, dbInstanceGoneStatusCodes) but NOT by NotFoundError,
//     which is exactly why the fix needs the explicit gone-code set plus a
//     follow-up DescribeDBInstance instead of relying on NotFoundError alone.
//   - A genuine 404 InvalidDBInstanceId.NotFound stays NotFoundError == true and is
//     NOT a gone-code match, keeping the recycle-bin (403) and not-found (404) paths
//     distinct.
//   - A non-terminal 403 (Throttling) matches neither, so a transient 403 is never
//     mistaken for a gone instance and remains retryable on mutation paths.
func TestRDSDBInstanceGoneStatusClassification(t *testing.T) {
	neutralMessage := "the request was not permitted in the current state"
	tests := []struct {
		name         string
		httpStatus   int
		code         string
		wantGone     bool
		wantNotFound bool
	}{
		{
			name:         "403 OperationDenied.DBInstanceStatus — refunded/recycle-bin parent",
			httpStatus:   403,
			code:         "OperationDenied.DBInstanceStatus",
			wantGone:     true,
			wantNotFound: false,
		},
		{
			name:         "403 OperationDenied.ReadDBInstanceStatus — read-side terminal lock",
			httpStatus:   403,
			code:         "OperationDenied.ReadDBInstanceStatus",
			wantGone:     true,
			wantNotFound: false,
		},
		{
			name:         "404 InvalidDBInstanceId.NotFound — real not-found",
			httpStatus:   404,
			code:         "InvalidDBInstanceId.NotFound",
			wantGone:     false,
			wantNotFound: true,
		},
		{
			name:         "403 Throttling — transient, not terminal",
			httpStatus:   403,
			code:         "Throttling",
			wantGone:     false,
			wantNotFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := newTestServerError(t, tt.httpStatus, tt.code, neutralMessage)
			assert.Equal(t, tt.wantGone, IsExpectedErrors(err, dbInstanceGoneStatusCodes),
				"IsExpectedErrors(err, dbInstanceGoneStatusCodes) for code %q", tt.code)
			assert.Equal(t, tt.wantNotFound, NotFoundError(err),
				"NotFoundError(err) for code %q", tt.code)
		})
	}
}

// TestRDSDBInstanceGoneStatusCodesContents guards the gone-code set itself: the two
// codes are load-bearing across DescribeDBInstance (choke point) and the
// child-resource Read/Delete follow-up checks. Drift here silently re-opens the
// refresh-hard-fail / destroy-hang regression, so any change must be deliberate.
func TestRDSDBInstanceGoneStatusCodesContents(t *testing.T) {
	assert.Equal(t,
		[]string{"OperationDenied.DBInstanceStatus", "OperationDenied.ReadDBInstanceStatus"},
		dbInstanceGoneStatusCodes)
	assert.NotContains(t, dbInstanceGoneStatusCodes, "OperationDenied.DBStatus",
		"DBStatus is a transient per-database lock, not a parent-gone signal — it must stay in OperationDeniedDBStatus (retryable), not here")
}
