// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test PaiWorkspace Config. >>> Resource test cases, automatically generated.
// Case TestConfig_0923_副本1730287211806 8577
func TestAccAliCloudPaiWorkspaceConfig_basic8577(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_pai_workspace_config.default"
	ra := resourceAttrInit(resourceId, AlicloudPaiWorkspaceConfigMap8577)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PaiWorkspaceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePaiWorkspaceConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPaiWorkspaceConfigBasicDependence8577)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"config_value": "\\\"[{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quota16a5gjqy2np\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"Products\\\\\\\":\\\\\\\"DLC,DSW\\\\\\\",\\\\\\\"SpecName\\\\\\\":\\\\\\\"aaa\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"8\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"4Gi\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\"}}]},{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quotabn1xxh0uc7j\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"SpecName\\\\\\\":\\\\\\\"testlinglan\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"1Gi\\\\\\\"},\\\\\\\"Products\\\\\\\":\\\\\\\"DSW\\\\\\\"}]}]\\\"",
					"workspace_id": "${alicloud_pai_workspace_workspace.workspace.id}",
					"labels": []map[string]interface{}{
						{
							"key":   "system.categoryName1",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName2",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName3",
							"value": "DSWAutoRecycle",
						},
					},
					"config_key":    "resourceSpecs",
					"category_name": "CommonQuotaConfig",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"config_value":  CHECKSET,
						"workspace_id":  CHECKSET,
						"labels.#":      "3",
						"config_key":    "resourceSpecs",
						"category_name": "CommonQuotaConfig",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config_value": "\\\"[{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quota16a5gjqy2np\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"Products\\\\\\\":\\\\\\\"DLC,DSW\\\\\\\",\\\\\\\"SpecName\\\\\\\":\\\\\\\"aaa\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"8\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"4Gi\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"0\\\\\\\"}}]},{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quotabn1xxh0uc7j\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"SpecName\\\\\\\":\\\\\\\"testlinglan\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"1Gi\\\\\\\"},\\\\\\\"Products\\\\\\\":\\\\\\\"DSW\\\\\\\"}]}]\\\"",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"config_value": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config_value": "\\\"[{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quota16a5gjqy2np\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"Products\\\\\\\":\\\\\\\"DLC,DSW\\\\\\\",\\\\\\\"SpecName\\\\\\\":\\\\\\\"aaa\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"8\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"4Gi\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\"}}]},{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quotabn1xxh0uc7j\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"SpecName\\\\\\\":\\\\\\\"testlinglan\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"1Gi\\\\\\\"},\\\\\\\"Products\\\\\\\":\\\\\\\"DSW\\\\\\\"}]}]\\\"",
					"labels": []map[string]interface{}{
						{
							"key":   "system.categoryName1",
							"value": "CommonResourceConfig",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"config_value": CHECKSET,
						"labels.#":     "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"labels": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"labels.#": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"labels": []map[string]interface{}{
						{
							"key":   "system.categoryName1",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName2",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName3",
							"value": "DSWAutoRecycle",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"labels.#": "3",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudPaiWorkspaceConfigMap8577 = map[string]string{}

func AlicloudPaiWorkspaceConfigBasicDependence8577(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_pai_workspace_workspace" "workspace" {
  description    = "函数引用： random(2,100)"
  display_name   = "test_pop_config_475"
  workspace_name = "test_pop_config_321"
  env_types      = ["prod"]
}


`, name)
}

// Case TestConfig 7451
func TestAccAliCloudPaiWorkspaceConfig_basic7451(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_pai_workspace_config.default"
	ra := resourceAttrInit(resourceId, AlicloudPaiWorkspaceConfigMap7451)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PaiWorkspaceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePaiWorkspaceConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPaiWorkspaceConfigBasicDependence7451)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"config_value": "\\\"[{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quota16a5gjqy2np\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"Products\\\\\\\":\\\\\\\"DLC,DSW\\\\\\\",\\\\\\\"SpecName\\\\\\\":\\\\\\\"aaa\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"8\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"4Gi\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\"}}]},{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quotabn1xxh0uc7j\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"SpecName\\\\\\\":\\\\\\\"testlinglan\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"1Gi\\\\\\\"},\\\\\\\"Products\\\\\\\":\\\\\\\"DSW\\\\\\\"}]}]\\\"",
					"workspace_id": "${alicloud_pai_workspace_workspace.workspace.id}",
					"labels": []map[string]interface{}{
						{
							"key":   "system.categoryName",
							"value": "CommonResourceConfig",
						},
						{
							"key":   "system.categoryName2",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName3",
							"value": "DSWAutoRecycle",
						},
					},
					"config_key":    "resourceSpecs",
					"category_name": "CommonQuotaConfig",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"config_value":  CHECKSET,
						"workspace_id":  CHECKSET,
						"labels.#":      "3",
						"config_key":    "resourceSpecs",
						"category_name": "CommonQuotaConfig",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config_value": "\\\"[{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quota16a5gjqy2np\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"Products\\\\\\\":\\\\\\\"DLC,DSW\\\\\\\",\\\\\\\"SpecName\\\\\\\":\\\\\\\"aaa\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"8\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"4Gi\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"0\\\\\\\"}}]},{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quotabn1xxh0uc7j\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"SpecName\\\\\\\":\\\\\\\"testlinglan\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"1Gi\\\\\\\"},\\\\\\\"Products\\\\\\\":\\\\\\\"DSW\\\\\\\"}]}]\\\"",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"config_value": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config_value": "\\\"[{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quota16a5gjqy2np\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"Products\\\\\\\":\\\\\\\"DLC,DSW\\\\\\\",\\\\\\\"SpecName\\\\\\\":\\\\\\\"aaa\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"8\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"4Gi\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\"}}]},{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quotabn1xxh0uc7j\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"SpecName\\\\\\\":\\\\\\\"testlinglan\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"1Gi\\\\\\\"},\\\\\\\"Products\\\\\\\":\\\\\\\"DSW\\\\\\\"}]}]\\\"",
					"labels": []map[string]interface{}{
						{
							"key":   "system.categoryName",
							"value": "CommonResourceConfig",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"config_value": CHECKSET,
						"labels.#":     "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"labels": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"labels.#": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudPaiWorkspaceConfigMap7451 = map[string]string{}

func AlicloudPaiWorkspaceConfigBasicDependence7451(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_pai_workspace_workspace" "workspace" {
  description    = "函数引用： random(2,100)"
  display_name   = "test_pop_config_567"
  workspace_name = "test_pop_config_101"
  env_types      = ["prod"]
}


`, name)
}

// Case TestConfig_0923 7950
func TestAccAliCloudPaiWorkspaceConfig_basic7950(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_pai_workspace_config.default"
	ra := resourceAttrInit(resourceId, AlicloudPaiWorkspaceConfigMap7950)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PaiWorkspaceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePaiWorkspaceConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPaiWorkspaceConfigBasicDependence7950)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"config_value": "\\\"[{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quota16a5gjqy2np\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"Products\\\\\\\":\\\\\\\"DLC,DSW\\\\\\\",\\\\\\\"SpecName\\\\\\\":\\\\\\\"aaa\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"8\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"4Gi\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\"}}]},{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quotabn1xxh0uc7j\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"SpecName\\\\\\\":\\\\\\\"testlinglan\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"1Gi\\\\\\\"},\\\\\\\"Products\\\\\\\":\\\\\\\"DSW\\\\\\\"}]}]\\\"",
					"workspace_id": "${alicloud_pai_workspace_workspace.workspace.id}",
					"labels": []map[string]interface{}{
						{
							"key":   "system.categoryName1",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName2",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName3",
							"value": "DSWAutoRecycle",
						},
					},
					"config_key":    "resourceSpecs",
					"category_name": "CommonQuotaConfig",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"config_value":  CHECKSET,
						"workspace_id":  CHECKSET,
						"labels.#":      "3",
						"config_key":    "resourceSpecs",
						"category_name": "CommonQuotaConfig",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config_value": "\\\"[{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quota16a5gjqy2np\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"Products\\\\\\\":\\\\\\\"DLC,DSW\\\\\\\",\\\\\\\"SpecName\\\\\\\":\\\\\\\"aaa\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"8\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"4Gi\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"0\\\\\\\"}}]},{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quotabn1xxh0uc7j\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"SpecName\\\\\\\":\\\\\\\"testlinglan\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"1Gi\\\\\\\"},\\\\\\\"Products\\\\\\\":\\\\\\\"DSW\\\\\\\"}]}]\\\"",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"config_value": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config_value": "\\\"[{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quota16a5gjqy2np\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"Products\\\\\\\":\\\\\\\"DLC,DSW\\\\\\\",\\\\\\\"SpecName\\\\\\\":\\\\\\\"aaa\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"8\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"4Gi\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\"}}]},{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quotabn1xxh0uc7j\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"SpecName\\\\\\\":\\\\\\\"testlinglan\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"1Gi\\\\\\\"},\\\\\\\"Products\\\\\\\":\\\\\\\"DSW\\\\\\\"}]}]\\\"",
					"labels": []map[string]interface{}{
						{
							"key":   "system.categoryName1",
							"value": "CommonResourceConfig",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"config_value": CHECKSET,
						"labels.#":     "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"labels": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"labels.#": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"labels": []map[string]interface{}{
						{
							"key":   "system.categoryName1",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName2",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName3",
							"value": "DSWAutoRecycle",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"labels.#": "3",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudPaiWorkspaceConfigMap7950 = map[string]string{}

func AlicloudPaiWorkspaceConfigBasicDependence7950(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_pai_workspace_workspace" "workspace" {
  description    = "函数引用： random(2,100)"
  display_name   = "test_pop_config_869"
  workspace_name = "test_pop_config_475"
  env_types      = ["prod"]
}


`, name)
}

// Case TestConfig_0923_副本1730287211806_副本1732087640194_副本1732087882610_副本1732170782019 9075
func TestAccAliCloudPaiWorkspaceConfig_basic9075(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_pai_workspace_config.default"
	ra := resourceAttrInit(resourceId, AlicloudPaiWorkspaceConfigMap9075)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PaiWorkspaceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePaiWorkspaceConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPaiWorkspaceConfigBasicDependence9075)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"config_value": "\\\"[{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quota16a5gjqy2np\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"Products\\\\\\\":\\\\\\\"DLC,DSW\\\\\\\",\\\\\\\"SpecName\\\\\\\":\\\\\\\"aaa\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"8\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"4Gi\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\"}}]},{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quotabn1xxh0uc7j\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"SpecName\\\\\\\":\\\\\\\"testlinglan\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"1Gi\\\\\\\"},\\\\\\\"Products\\\\\\\":\\\\\\\"DSW\\\\\\\"}]}]\\\"",
					"workspace_id": "${alicloud_pai_workspace_workspace.workspace.id}",
					"labels": []map[string]interface{}{
						{
							"key":   "system.categoryName1",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName2",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName3",
							"value": "DSWAutoRecycle",
						},
					},
					"config_key":    "resourceSpecs",
					"category_name": "CommonQuotaConfig",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"config_value":  CHECKSET,
						"workspace_id":  CHECKSET,
						"labels.#":      "3",
						"config_key":    "resourceSpecs",
						"category_name": "CommonQuotaConfig",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config_value": "\\\"[{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quota16a5gjqy2np\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"Products\\\\\\\":\\\\\\\"DLC,DSW\\\\\\\",\\\\\\\"SpecName\\\\\\\":\\\\\\\"aaa\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"8\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"4Gi\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"0\\\\\\\"}}]},{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quotabn1xxh0uc7j\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"SpecName\\\\\\\":\\\\\\\"testlinglan\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"1Gi\\\\\\\"},\\\\\\\"Products\\\\\\\":\\\\\\\"DSW\\\\\\\"}]}]\\\"",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"config_value": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config_value": "\\\"[{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quota16a5gjqy2np\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"Products\\\\\\\":\\\\\\\"DLC,DSW\\\\\\\",\\\\\\\"SpecName\\\\\\\":\\\\\\\"aaa\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"8\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"4Gi\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\"}}]},{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quotabn1xxh0uc7j\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"SpecName\\\\\\\":\\\\\\\"testlinglan\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"1Gi\\\\\\\"},\\\\\\\"Products\\\\\\\":\\\\\\\"DSW\\\\\\\"}]}]\\\"",
					"labels": []map[string]interface{}{
						{
							"key":   "system.categoryName1",
							"value": "CommonResourceConfig",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"config_value": CHECKSET,
						"labels.#":     "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"labels": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"labels.#": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"labels": []map[string]interface{}{
						{
							"key":   "system.categoryName1",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName2",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName3",
							"value": "DSWAutoRecycle",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"labels.#": "3",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudPaiWorkspaceConfigMap9075 = map[string]string{}

func AlicloudPaiWorkspaceConfigBasicDependence9075(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_pai_workspace_workspace" "workspace" {
  description    = "82"
  display_name   = "test_pop_config_407"
  workspace_name = "test_pop_config_310"
  env_types      = ["prod"]
}


`, name)
}

// Case TestConfig_0923_副本1730287211806_副本1732087640194_副本1732087882610_副本1732170782019_副本1740729226197 10420
func TestAccAliCloudPaiWorkspaceConfig_basic10420(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_pai_workspace_config.default"
	ra := resourceAttrInit(resourceId, AlicloudPaiWorkspaceConfigMap10420)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PaiWorkspaceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePaiWorkspaceConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPaiWorkspaceConfigBasicDependence10420)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"config_value": "\\\"[{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quota16a5gjqy2np\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"Products\\\\\\\":\\\\\\\"DLC,DSW\\\\\\\",\\\\\\\"SpecName\\\\\\\":\\\\\\\"aaa\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"8\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"4Gi\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\"}}]},{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quotabn1xxh0uc7j\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"SpecName\\\\\\\":\\\\\\\"testlinglan\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"1Gi\\\\\\\"},\\\\\\\"Products\\\\\\\":\\\\\\\"DSW\\\\\\\"}]}]\\\"",
					"workspace_id": "${alicloud_pai_workspace_workspace.workspace.id}",
					"labels": []map[string]interface{}{
						{
							"key":   "system.categoryName1",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName2",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName3",
							"value": "DSWAutoRecycle",
						},
					},
					"config_key":    "resourceSpecs",
					"category_name": "CommonQuotaConfig",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"config_value":  CHECKSET,
						"workspace_id":  CHECKSET,
						"labels.#":      "3",
						"config_key":    "resourceSpecs",
						"category_name": "CommonQuotaConfig",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config_value": "\\\"[{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quota16a5gjqy2np\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"Products\\\\\\\":\\\\\\\"DLC,DSW\\\\\\\",\\\\\\\"SpecName\\\\\\\":\\\\\\\"aaa\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"8\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"4Gi\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"0\\\\\\\"}}]},{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quotabn1xxh0uc7j\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"SpecName\\\\\\\":\\\\\\\"testlinglan\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"1Gi\\\\\\\"},\\\\\\\"Products\\\\\\\":\\\\\\\"DSW\\\\\\\"}]}]\\\"",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"config_value": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config_value": "\\\"[{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quota16a5gjqy2np\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"Products\\\\\\\":\\\\\\\"DLC,DSW\\\\\\\",\\\\\\\"SpecName\\\\\\\":\\\\\\\"aaa\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"8\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"4Gi\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\"}}]},{\\\\\\\"Type\\\\\\\":\\\\\\\"Quota\\\\\\\",\\\\\\\"Key\\\\\\\":\\\\\\\"quotabn1xxh0uc7j\\\\\\\",\\\\\\\"Value\\\\\\\":[{\\\\\\\"SpecName\\\\\\\":\\\\\\\"testlinglan\\\\\\\",\\\\\\\"Spec\\\\\\\":{\\\\\\\"CPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"GPU\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"Memory\\\\\\\":\\\\\\\"1Gi\\\\\\\"},\\\\\\\"Products\\\\\\\":\\\\\\\"DSW\\\\\\\"}]}]\\\"",
					"labels": []map[string]interface{}{
						{
							"key":   "system.categoryName1",
							"value": "CommonResourceConfig",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"config_value": CHECKSET,
						"labels.#":     "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"labels": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"labels.#": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"labels": []map[string]interface{}{
						{
							"key":   "system.categoryName1",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName2",
							"value": "DSWAutoRecycle",
						},
						{
							"key":   "system.categoryName3",
							"value": "DSWAutoRecycle",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"labels.#": "3",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudPaiWorkspaceConfigMap10420 = map[string]string{}

func AlicloudPaiWorkspaceConfigBasicDependence10420(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_pai_workspace_workspace" "workspace" {
  description    = "294"
  display_name   = "test_pop_config_762"
  workspace_name = "test_pop_config_830"
  env_types      = ["prod"]
}


`, name)
}

// Test PaiWorkspace Config. <<< Resource test cases, automatically generated.
