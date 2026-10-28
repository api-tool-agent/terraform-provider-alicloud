package cms

import "github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"

// NewCmsServiceV2 returns a CmsServiceV2 bound to the given client. It exists so
// package-external acceptance tests can construct the service for destroy checks.
func NewCmsServiceV2(client *connectivity.AliyunClient) *CmsServiceV2 {
	return &CmsServiceV2{client: client}
}
