// Package cms. This file is generated automatically. Please do not modify it manually, thank you!
package cms

import (
	tferrs "github.com/aliyun/terraform-provider-alicloud/alicloud/errs"
	tflogs "github.com/aliyun/terraform-provider-alicloud/alicloud/logs"

	tfconvert "github.com/aliyun/terraform-provider-alicloud/alicloud/convert"

	tfhelper "github.com/aliyun/terraform-provider-alicloud/alicloud/helper"

	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	tfretry "github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/tidwall/sjson"
)

type CmsServiceV2 struct {
	client *connectivity.AliyunClient
}

// DescribeCmsAlertRuleV2 <<< Encapsulated get interface for Cms AlertRuleV2.

func (s *CmsServiceV2) DescribeCmsAlertRuleV2(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var body map[string]interface{}
	request = make(map[string]interface{})
	query = make(map[string]*string)
	query["RegionId"] = tfhelper.StringPointer(client.RegionId)
	jsonString := tfconvert.ConvertObjectToJsonString(request)
	jsonString, _ = sjson.Set(jsonString, "filter.uuid.eq", id)
	_ = json.Unmarshal([]byte(jsonString), &request)
	body = request

	action := fmt.Sprintf("/queryAlertRules")

	wait := tfhelper.IncrementalWait(3*time.Second, 5*time.Second)
	err = tfretry.Retry(1*time.Minute, func() *tfretry.RetryError {
		response, err = client.RoaPost("Cms", "2024-03-30", action, query, nil, body, true)

		if err != nil {
			if tfhelper.NeedRetry(err) {
				wait()
				return tfretry.RetryableError(err)
			}
			return tfretry.NonRetryableError(err)
		}
		return nil
	})
	tflogs.AddDebug(action, response, request)
	if err != nil {
		if tferrs.IsExpectedErrors(err, []string{"ResourceNotFound"}) {
			return object, tferrs.WrapErrorf(tferrs.NotFoundErr("AlertRuleV2", id), tferrs.NotFoundMsg, response)
		}
		return object, tferrs.WrapErrorf(err, tferrs.DefaultErrorMsg, id, action, tferrs.AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.data.alertRules[*]", response)
	if err != nil {
		return object, tferrs.WrapErrorf(err, tferrs.FailedGetAttributeMsg, id, "$.data.alertRules[*]", response)
	}

	if len(v.([]interface{})) == 0 {
		return object, tferrs.WrapErrorf(tferrs.NotFoundErr("AlertRuleV2", id), tferrs.NotFoundMsg, response)
	}

	return v.([]interface{})[0].(map[string]interface{}), nil
}

func (s *CmsServiceV2) CmsAlertRuleV2StateRefreshFunc(id string, field string, failStates []string) tfretry.StateRefreshFunc {
	return s.CmsAlertRuleV2StateRefreshFuncWithApi(id, field, failStates, s.DescribeCmsAlertRuleV2)
}

func (s *CmsServiceV2) CmsAlertRuleV2StateRefreshFuncWithApi(id string, field string, failStates []string, call func(id string) (map[string]interface{}, error)) tfretry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := call(id)
		if err != nil {
			if tferrs.NotFoundError(err) {
				return object, "", nil
			}
			return nil, "", tferrs.WrapError(err)
		}
		v, err := jsonpath.Get(field, object)
		currentStatus := fmt.Sprint(v)

		if strings.HasPrefix(field, "#") {
			v, _ := jsonpath.Get(strings.TrimPrefix(field, "#"), object)
			if v != nil {
				currentStatus = "#CHECKSET"
			}
		}

		for _, failState := range failStates {
			if currentStatus == failState {
				return object, currentStatus, tferrs.WrapError(tferrs.Error(tferrs.FailedToReachTargetStatus, currentStatus))
			}
		}
		return object, currentStatus, nil
	}
}

// DescribeCmsAlertRuleV2 >>> Encapsulated.
