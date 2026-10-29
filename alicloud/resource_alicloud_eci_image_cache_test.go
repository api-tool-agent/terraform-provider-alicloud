package alicloud

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/eci"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func init() {
	resource.AddTestSweepers(
		"alicloud_eci_image_cache",
		&resource.Sweeper{
			Name: "alicloud_eci_image_cache",
			F:    testSweepEciImageCache,
		})
}

func testSweepEciImageCache(region string) error {
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		return WrapError(err)
	}
	client := rawClient.(*connectivity.AliyunClient)
	queryRequest := eci.CreateDescribeImageCachesRequest()
	var allCaches []eci.DescribeImageCachesImageCache0

	raw, err := client.WithEciClient(func(eciClient *eci.Client) (interface{}, error) {
		return eciClient.DescribeImageCaches(queryRequest)
	})
	if err != nil {
		log.Printf("[ERROR] %s get an error %#v", queryRequest.GetActionName(), err)
	}
	addDebug(queryRequest.GetActionName(), raw)
	response, _ := raw.(*eci.DescribeImageCachesResponse)
	for _, cache := range response.ImageCaches {
		if strings.HasPrefix(cache.ImageCacheName, "tf-testacc") {
			allCaches = append(allCaches, cache)
		} else {
			log.Printf("Skip %#v", cache)
		}
	}

	removeRequest := eci.CreateDeleteImageCacheRequest()
	removeRequest.ImageCacheId = ""
	for _, cache := range allCaches {
		removeRequest.ImageCacheId = cache.ImageCacheId
		raw, err := client.WithEciClient(func(eciClient *eci.Client) (interface{}, error) {
			return eciClient.DeleteImageCache(removeRequest)
		})
		if err != nil {
			log.Printf("[ERROR] %s get an error %s", removeRequest.GetActionName(), err)
		}
		addDebug(removeRequest.GetActionName(), raw)
	}
	return nil
}

func TestAccAliCloudEciImageCache_basic(t *testing.T) {
	var v eci.DescribeImageCachesImageCache0
	checkoutSupportedRegions(t, true, connectivity.EciContainerGroupRegions)
	resourceId := "alicloud_eci_image_cache.default"
	ra := resourceAttrInit(resourceId, EciOpenapiImageCacheMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EciService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEciImageCache")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1000000, 9999999)
	name := fmt.Sprintf("tf-testAccEciImageCache%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, EciOpenapiImageCacheBasicdependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		ProviderFactories: testAccProviderFactory,
		CheckDestroy:      rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"image_cache_name":       strings.ToLower(name),
					"images":                 []string{fmt.Sprintf("registry-vpc.%s.aliyuncs.com/eci_open/nginx:alpine", defaultRegionToTest)},
					"security_group_id":      "${alicloud_security_group.group.id}",
					"vswitch_id":             "${data.alicloud_vpcs.default.vpcs.0.vswitch_ids.0}",
					"eip_instance_id":        "${alicloud_eip_address.default.id}",
					"resource_group_id":      "${data.alicloud_resource_manager_resource_groups.default.groups.0.id}",
					"image_cache_size":       "20",
					"retention_days":         "1",
					"zone_id":                "${data.alicloud_vswitches.default.vswitches.0.zone_id}",
					"auto_match_image_cache": "false",
					"flash":                  "false",
					"flash_copy_count":       "1",
					"standard_copy_count":    "1",
					"elimination_strategy":   "LRU",
					"os_type":                "Linux",
					"insecure_registry":      "harbor.example.com",
					"plain_http_registry":    "harbor.example.com",
					"depends_on ":            []string{"alicloud_eip_address.default"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"image_cache_name":       strings.ToLower(name),
						"images.#":               "1",
						"resource_group_id":      CHECKSET,
						"elimination_strategy":   "LRU",
						"os_type":                "Linux",
						"insecure_registry":      "harbor.example.com",
						"plain_http_registry":    "harbor.example.com",
						"auto_match_image_cache": "false",
						"flash":                  "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"image_cache_name":     strings.ToLower(name),
					"images":               []string{fmt.Sprintf("registry-vpc.%s.aliyuncs.com/eci_open/nginx:alpine", defaultRegionToTest)},
					"security_group_id":    "${alicloud_security_group.group.id}",
					"vswitch_id":           "${data.alicloud_vpcs.default.vpcs.0.vswitch_ids.0}",
					"eip_instance_id":      "${alicloud_eip_address.default.id}",
					"resource_group_id":    "${data.alicloud_resource_manager_resource_groups.default.groups.0.id}",
					"image_cache_size":     "20",
					"retention_days":       "1",
					"zone_id":              "${data.alicloud_vswitches.default.vswitches.0.zone_id}",
					"flash_copy_count":     "1",
					"standard_copy_count":  "1",
					"elimination_strategy": "LRU",
					"os_type":              "Linux",
					"insecure_registry":    "harbor.example.com",
					"plain_http_registry":  "harbor.example.com",
					"acr_registry_info": []map[string]interface{}{
						{
							"arn_service":   "acs:ram::1234567890:role/acr-test",
							"arn_user":      "acs:ram::1234567890:role/acr-test-user",
							"domain":        []string{fmt.Sprintf("registry.%s.cr.aliyuncs.com", defaultRegionToTest)},
							"instance_id":   "acr-1",
							"instance_name": "acr_test_name",
							"region_id":     defaultRegionToTest,
						},
					},
					"image_registry_credential": []map[string]interface{}{
						{
							"server":    fmt.Sprintf("registry-vpc.%s.aliyuncs.com", defaultRegionToTest),
							"user_name": name,
							"password":  "tftestacc",
						},
					},
					"depends_on ": []string{"alicloud_eip_address.default"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"image_cache_name":                   strings.ToLower(name),
						"images.#":                           "1",
						"acr_registry_info.#":                "1",
						"acr_registry_info.0.instance_id":    "acr-1",
						"acr_registry_info.0.arn_service":    "acs:ram::1234567890:role/acr-test",
						"acr_registry_info.0.arn_user":       "acs:ram::1234567890:role/acr-test-user",
						"image_registry_credential.#":        "1",
						"image_registry_credential.0.server": fmt.Sprintf("registry-vpc.%s.aliyuncs.com", defaultRegionToTest),
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_group_id", "vswitch_id", "resource_group_id", "eip_instance_id", "image_cache_size", "retention_days", "zone_id", "elimination_strategy", "os_type", "insecure_registry", "plain_http_registry", "flash_copy_count", "standard_copy_count", "acr_registry_info", "image_registry_credential"},
			},
		},
	})
}

var EciOpenapiImageCacheMap = map[string]string{
	"container_group_id": CHECKSET,
	"status":             CHECKSET,
	"eip_instance_id":    CHECKSET,
}

func EciOpenapiImageCacheBasicdependence(name string) string {
	return fmt.Sprintf(`
	data "alicloud_vpcs" "default" {
	  name_regex = "^default-NODELETING$"
	}
	data "alicloud_vswitches" "default" {
	  ids = [data.alicloud_vpcs.default.vpcs.0.vswitch_ids.0]
	}
	resource "alicloud_security_group" "group" {
	  name        = "%[1]s"
	  description = "tf-eci-image-test"
	  vpc_id      = data.alicloud_vpcs.default.vpcs.0.id
	}
	resource "alicloud_eip_address" "default" {
	  address_name = "%[1]s"
	}
	data "alicloud_resource_manager_resource_groups" "default" {
	  name_regex = "default"
	}
`, name)
}
