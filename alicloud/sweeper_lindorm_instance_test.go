package alicloud

// The Lindorm instance sweeper stays in the root test binary: resource.AddTestSweepers
// registration must happen in package alicloud, alongside the other sweepers.

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

func init() {
	resource.AddTestSweepers("alicloud_lindorm_instance", &resource.Sweeper{
		Name: "alicloud_lindorm_instance",
		F:    testSweepLindormInstances,
	})
}

func testSweepLindormInstances(region string) error {
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		return fmt.Errorf("error getting AliCloud client: %s", err)
	}
	client := rawClient.(*connectivity.AliyunClient)
	prefixes := []string{
		"tf-testAcc",
		"tf_testAcc",
	}
	action := "GetLindormInstanceList"
	request := make(map[string]interface{})
	request["PageSize"] = PageSizeLarge
	request["PageNumber"] = 1
	var response map[string]interface{}
	lindormInstanceIds := make([]string, 0)
	for {
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = retry.Retry(5*time.Minute, func() *retry.RetryError {
			response, err = client.RpcPost("hitsdb", "2020-06-15", action, nil, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return retry.RetryableError(err)
				}
				return retry.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DataDefaultErrorMsg, "alicloud_lindorm_instances", action, AlibabaCloudSdkGoERROR)
		}
		resp, err := jsonpath.Get("$.InstanceList", response)
		if err != nil {
			return WrapErrorf(err, FailedGetAttributeMsg, action, "$.InstanceList", response)
		}
		result, _ := resp.([]interface{})
		for _, v := range result {
			skip := true
			item := v.(map[string]interface{})
			if !sweepAll() {
				for _, prefix := range prefixes {
					if strings.HasPrefix(strings.ToLower(fmt.Sprint(item["InstanceAlias"])), strings.ToLower(prefix)) {
						skip = false
						break
					}
				}
				if skip {
					log.Printf("[INFO] Skipping Lindorm Instance: %v (%v)", item["InstanceAlias"], item["InstanceId"])
					continue
				}
			}
			lindormInstanceIds = append(lindormInstanceIds, fmt.Sprint(item["InstanceId"]))
		}
		if len(result) < PageSizeLarge {
			break
		}
		request["PageNumber"] = request["PageNumber"].(int) + 1
	}

	for _, id := range lindormInstanceIds {
		log.Printf("[INFO] Deleting Lindorm Instance: %s", id)
		action := "ReleaseLindormInstance"
		request := map[string]interface{}{
			"InstanceId": id,
		}
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = retry.Retry(1*time.Minute, func() *retry.RetryError {
			_, err = client.RpcPost("hitsdb", "2020-06-15", action, nil, request, false)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return retry.RetryableError(err)
				}
				return retry.NonRetryableError(err)
			}
			return nil
		})
		if err != nil {
			log.Printf("[ERROR] Failed to delete lindorm Instance (%s): %s", id, err)
		}
	}
	return nil
}
