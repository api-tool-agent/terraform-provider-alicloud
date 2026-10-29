package polardb

import (
	"fmt"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/errs"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/helper"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/logs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

// Service wraps the PolarDB API calls used by the resources and data sources
// of this package.
type Service struct {
	client *connectivity.AliyunClient
}

// NewService is the exported constructor: the client field stays unexported.
func NewService(client *connectivity.AliyunClient) *Service {
	return &Service{client: client}
}

func (s *Service) rpcPost(action string, request map[string]interface{}, timeout time.Duration) (map[string]interface{}, error) {
	var response map[string]interface{}
	var err error
	wait := helper.IncrementalWait(3*time.Second, 3*time.Second)
	err = retry.Retry(timeout, func() *retry.RetryError {
		response, err = s.client.RpcPost("polardb", "2017-08-01", action, nil, request, false)
		if err != nil {
			if helper.NeedRetry(err) {
				wait()
				return retry.RetryableError(err)
			}
			return retry.NonRetryableError(err)
		}
		return nil
	})
	logs.AddDebug(action, response, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// DescribeAIDBClusterDataset returns one dataset of an AI cluster, identified by
// its own ID. The DescribeAIDBClusterDatasets API only offers a list entry, so
// the matching record is selected from the returned page.
func (s *Service) DescribeAIDBClusterDataset(dbClusterId, datasetId string) (map[string]interface{}, error) {
	request := map[string]interface{}{
		"DBClusterId": dbClusterId,
		"DatasetId":   datasetId,
	}
	if s.client.RegionId != "" {
		request["RegionId"] = s.client.RegionId
	}
	response, err := s.rpcPost("DescribeAIDBClusterDatasets", request, 5*time.Minute)
	if err != nil {
		return nil, errs.WrapError(err)
	}
	if datasets, ok := response["Datasets"].([]interface{}); ok {
		for _, item := range datasets {
			object, ok := item.(map[string]interface{})
			if !ok {
				continue
			}
			if fmt.Sprint(object["DatasetId"]) == datasetId {
				return object, nil
			}
		}
	}
	return nil, errs.WrapErrorf(errs.NotFoundErr("PolarDB", datasetId), errs.NotFoundWithResponse, response)
}
