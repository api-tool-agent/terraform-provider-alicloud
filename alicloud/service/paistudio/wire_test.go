package paistudio

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// These TestUnits pin the request/response wire that the production CRUD path
// builds and then sends through AliyunClient.RoaPost/RoaGet/RoaPut/RoaDelete:
// the ROA path templates, the JSON body shape sent for the form-data
// CreateAlgorithm/UpdateAlgorithm parameters, the GetAlgorithm state mapping,
// the typed NoSuchObject not-found detection and the GetAlgorithm identity
// guard. They are offline: no client, no network, no cloud call.

func TestUnitPaiStudioAlgorithmPath(t *testing.T) {
	if got := algorithmPath(""); got != "/api/v1/algorithms" {
		t.Errorf("create path = %q, want /api/v1/algorithms", got)
	}
	if got := algorithmPath("algo-123"); got != "/api/v1/algorithms/algo-123" {
		t.Errorf("object path = %q, want /api/v1/algorithms/algo-123", got)
	}
}

func TestUnitPaiStudioAlgorithmCreateBody(t *testing.T) {
	body := algorithmCreateBody(algorithmResourceModel{
		AlgorithmName:        types.StringValue("tf-testAccPaistudioAlgorithm-1"),
		WorkspaceId:          types.StringValue("ws-1"),
		DisplayName:          types.StringNull(),
		AlgorithmDescription: types.StringValue("desc"),
	})
	if body["AlgorithmName"] != "tf-testAccPaistudioAlgorithm-1" {
		t.Errorf("AlgorithmName = %v, want tf-testAccPaistudioAlgorithm-1", body["AlgorithmName"])
	}
	if body["WorkspaceId"] != "ws-1" {
		t.Errorf("WorkspaceId = %v, want ws-1", body["WorkspaceId"])
	}
	if body["AlgorithmDescription"] != "desc" {
		t.Errorf("AlgorithmDescription = %v, want desc", body["AlgorithmDescription"])
	}
	if _, ok := body["DisplayName"]; ok {
		t.Errorf("null DisplayName must not be sent, got %v", body["DisplayName"])
	}

	// Unknown and unset optional values are dropped, never sent as "".
	body = algorithmCreateBody(algorithmResourceModel{
		AlgorithmName:        types.StringValue("tf-testAccPaistudioAlgorithm-2"),
		WorkspaceId:          types.StringUnknown(),
		DisplayName:          types.StringNull(),
		AlgorithmDescription: types.StringNull(),
	})
	if len(body) != 1 {
		t.Errorf("create body with only the name should carry one key, got %#v", body)
	}
}

func TestUnitPaiStudioAlgorithmUpdateBody(t *testing.T) {
	// The UpdateAlgorithm contract only accepts DisplayName and
	// AlgorithmDescription; no other key may leak into the request body.
	body := algorithmUpdateBody(algorithmResourceModel{
		Id:                   types.StringValue("algo-1"),
		AlgorithmName:        types.StringValue("should-not-be-sent"),
		WorkspaceId:          types.StringValue("ws-1"),
		DisplayName:          types.StringValue("display-2"),
		AlgorithmDescription: types.StringValue("desc-2"),
	})
	if len(body) != 2 {
		t.Errorf("update body should carry exactly the two updatable keys, got %#v", body)
	}
	if body["DisplayName"] != "display-2" || body["AlgorithmDescription"] != "desc-2" {
		t.Errorf("update body = %#v, want DisplayName=display-2, AlgorithmDescription=desc-2", body)
	}

	body = algorithmUpdateBody(algorithmResourceModel{
		DisplayName:          types.StringNull(),
		AlgorithmDescription: types.StringUnknown(),
	})
	if len(body) != 0 {
		t.Errorf("update body with no set values must be empty, got %#v", body)
	}
}

func TestUnitPaiStudioAlgorithmStateMapping(t *testing.T) {
	object := map[string]interface{}{
		"AlgorithmId":          "algo-1",
		"AlgorithmName":        "tf-testAccPaistudioAlgorithm-1",
		"WorkspaceId":          "ws-1",
		"DisplayName":          "display-1",
		"AlgorithmDescription": "desc-1",
		"AlgorithmProvider":    "system",
		"UserId":               "12345",
		"TenantId":             "67890",
		"GmtCreateTime":        "2026-10-18T00:00:00Z",
		"GmtModifiedTime":      "2026-10-18T01:00:00Z",
	}
	state := algorithmState(object)
	if state.Id.ValueString() != "algo-1" {
		t.Errorf("Id = %q, want algo-1", state.Id.ValueString())
	}

	// Check each mapped field explicitly.
	checks := []struct {
		got  types.String
		want string
	}{
		{state.AlgorithmName, "tf-testAccPaistudioAlgorithm-1"},
		{state.WorkspaceId, "ws-1"},
		{state.DisplayName, "display-1"},
		{state.AlgorithmDescription, "desc-1"},
		{state.AlgorithmProvider, "system"},
		{state.UserId, "12345"},
		{state.TenantId, "67890"},
		{state.GmtCreateTime, "2026-10-18T00:00:00Z"},
		{state.GmtModifiedTime, "2026-10-18T01:00:00Z"},
	}
	for i, c := range checks {
		if c.got.ValueString() != c.want {
			t.Errorf("field %d = %q, want %q", i, c.got.ValueString(), c.want)
		}
	}

	// Empty strings and absent keys map to null, not "".
	empty := algorithmState(map[string]interface{}{
		"AlgorithmId":       "algo-2",
		"AlgorithmName":     "n",
		"AlgorithmProvider": "",
	})
	if empty.Id.ValueString() != "algo-2" {
		t.Errorf("Id = %q, want algo-2", empty.Id.ValueString())
	}
	for i, s := range []types.String{
		empty.WorkspaceId, empty.DisplayName, empty.AlgorithmDescription,
		empty.AlgorithmProvider, empty.UserId, empty.TenantId,
		empty.GmtCreateTime, empty.GmtModifiedTime,
	} {
		if !s.IsNull() {
			t.Errorf("empty/absent field %d must map to null, got %q", i, s.ValueString())
		}
	}
}

func TestUnitPaiStudioAlgorithmNotExistDetection(t *testing.T) {
	if isAlgorithmNotExist(nil) {
		t.Error("nil error must not be treated as not-found")
	}
	if isAlgorithmNotExist(errors.New("connection reset")) {
		t.Error("a plain error must not be treated as not-found")
	}
	if isAlgorithmNotExist(&tea.SDKError{Code: tea.String("Forbidden")}) {
		t.Error("a different SDK error code must not be treated as not-found")
	}
	if !isAlgorithmNotExist(&tea.SDKError{Code: tea.String("NoSuchObject")}) {
		t.Error("SDK error NoSuchObject must be treated as not-found")
	}
	// A message merely containing the code must not be misread.
	if isAlgorithmNotExist(fmt.Errorf("server said NoSuchObject in text")) {
		t.Error("a plain error mentioning NoSuchObject must not be treated as not-found")
	}
}

func TestUnitPaiStudioAlgorithmValidateResponse(t *testing.T) {
	if err := validateAlgorithmResponse(nil, "algo-1"); err == nil {
		t.Error("nil response must be rejected")
	}
	if err := validateAlgorithmResponse(map[string]interface{}{}, "algo-1"); err == nil {
		t.Error("missing AlgorithmId must be rejected")
	}
	if err := validateAlgorithmResponse(map[string]interface{}{"AlgorithmId": 42}, "algo-1"); err == nil {
		t.Error("wrong-typed AlgorithmId must be rejected")
	}
	if err := validateAlgorithmResponse(map[string]interface{}{"AlgorithmId": "other"}, "algo-1"); err == nil {
		t.Error("mismatched AlgorithmId must be rejected")
	}
	if err := validateAlgorithmResponse(map[string]interface{}{"AlgorithmId": "algo-1"}, "algo-1"); err != nil {
		t.Errorf("matching response must validate, got %s", err)
	}
}
