package alicloud

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccAliCloudRdsParameterGroup_basic(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_parameter_group.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsParameterGroupMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsParameterGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf_testAccAlicloudRdsParameterGroup%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsParameterGroupBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		ProviderFactories: testAccProviderFactory,
		CheckDestroy:      rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"engine":         "mysql",
					"engine_version": `5.7`,
					"param_detail": []map[string]interface{}{
						{
							"param_name":  "back_log",
							"param_value": `3000`,
						},
						{
							"param_name":  "wait_timeout",
							"param_value": `86400`,
						},
					},
					"parameter_group_name": "${var.name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"engine":                 "mysql",
						"engine_version":         "5.7",
						"param_detail.#":         "2",
						"parameter_group_name":   name,
						"related_custins_info.#": "0",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"param_detail": []map[string]interface{}{
						{
							"param_name":  "back_log",
							"param_value": `4000`,
						},
						{
							"param_name":  "wait_timeout",
							"param_value": `86460`,
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"param_detail.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"parameter_group_desc": "update_test",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"parameter_group_desc": "update_test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"parameter_group_name": name + "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"parameter_group_name": name + "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"param_detail": []map[string]interface{}{
						{
							"param_name":  "back_log",
							"param_value": `3000`,
						},
						{
							"param_name":  "wait_timeout",
							"param_value": `86400`,
						},
					},
					"parameter_group_desc": "test",
					"parameter_group_name": "${var.name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"param_detail.#":       "2",
						"parameter_group_desc": "test",
						"parameter_group_name": name,
					}),
				),
			},
		},
	})
}

func TestAccAliCloudRdsParameterGroup_basic1(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_parameter_group.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsParameterGroupMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsParameterGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf_testAccAlicloudRdsParameterGroup%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsParameterGroupBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		ProviderFactories: testAccProviderFactory,
		CheckDestroy:      rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"engine":         "mysql",
					"engine_version": `5.7`,
					"param_detail": []map[string]interface{}{
						{
							"param_name":  "back_log",
							"param_value": `3000`,
						},
						{
							"param_name":  "wait_timeout",
							"param_value": `86400`,
						},
					},
					"parameter_group_name": "${var.name}",
					"parameter_group_desc": "update_test",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"engine":               "mysql",
						"engine_version":       "5.7",
						"param_detail.#":       "2",
						"parameter_group_name": name,
						"parameter_group_desc": "update_test",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

var AlicloudRdsParameterGroupMap0 = map[string]string{}

func AlicloudRdsParameterGroupBasicDependence0(name string) string {
	return fmt.Sprintf(`
variable "name" {
			default = "%s"
		}
`, name)
}

// lintignore: R001
func TestUnitAlicloudRdsParameterGroup(t *testing.T) {
	p := Provider().ResourcesMap
	dInit, _ := schema.InternalMap(p["alicloud_rds_parameter_group"].Schema).Data(nil, nil)
	dExisted, _ := schema.InternalMap(p["alicloud_rds_parameter_group"].Schema).Data(nil, nil)
	dInit.MarkNewResource()
	attributes := map[string]interface{}{
		"engine":         "CreateParameterGroupValue",
		"engine_version": "CreateParameterGroupValue",
		"param_detail": []map[string]interface{}{
			{
				"param_name":  "CreateParameterGroupValue",
				"param_value": "CreateParameterGroupValue",
			},
		},
		"parameter_group_name": "CreateParameterGroupValue",
		"parameter_group_desc": "CreateParameterGroupValue",
	}
	for key, value := range attributes {
		err := dInit.Set(key, value)
		assert.Nil(t, err)
		err = dExisted.Set(key, value)
		assert.Nil(t, err)
		if err != nil {
			log.Printf("[ERROR] the field %s setting error", key)
		}
	}
	region := os.Getenv("ALICLOUD_REGION")
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		t.Skipf("Skipping the test case with err: %s", err)
		t.Skipped()
	}
	rawClient = rawClient.(*connectivity.AliyunClient)
	ReadMockResponse := map[string]interface{}{
		// DescribeParameterGroup
		"ParamGroup": map[string]interface{}{
			"ParameterGroup": []interface{}{
				map[string]interface{}{
					"ParameterGroupId": "CreateParameterGroupValue",
					"Engine":           "CreateParameterGroupValue",
					"EngineVersion":    "CreateParameterGroupValue",
					"ParamDetail": map[string]interface{}{
						"ParameterDetail": []interface{}{
							map[string]interface{}{
								"ParamName":  "CreateParameterGroupValue",
								"ParamValue": "CreateParameterGroupValue",
							},
						},
					},
					"ParameterGroupName": "CreateParameterGroupValue",
					"ParameterGroupDesc": "CreateParameterGroupValue",
				},
			},
		},
		"ParameterGroupId": "CreateParameterGroupValue",
		"RelatedCustinsInfo": map[string]interface{}{
			"RelatedCustinsInfo": []interface{}{
				map[string]interface{}{
					"DBInstanceName": "CreateParameterGroupValue",
					"AppliedTime":    "CreateParameterGroupValue",
				},
			},
		},
	}
	CreateMockResponse := map[string]interface{}{
		// CreateParameterGroup
		"ParameterGroupId": "CreateParameterGroupValue",
	}
	failedResponseMock := func(errorCode string) (map[string]interface{}, error) {
		return nil, &tea.SDKError{
			Code:       String(errorCode),
			Data:       String(errorCode),
			Message:    String(errorCode),
			StatusCode: tea.Int(400),
		}
	}
	notFoundResponseMock := func(errorCode string) (map[string]interface{}, error) {
		return nil, GetNotFoundErrorFromString(GetNotFoundMessage("alicloud_rds_parameter_group", errorCode))
	}
	successResponseMock := func(operationMockResponse map[string]interface{}) (map[string]interface{}, error) {
		if len(operationMockResponse) > 0 {
			mapMerge(ReadMockResponse, operationMockResponse)
		}
		return ReadMockResponse, nil
	}

	// Create
	patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewRdsClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAlicloudRdsParameterGroupCreate(dInit, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	ReadMockResponseDiff := map[string]interface{}{
		// DescribeParameterGroup Response
		"ParameterGroupId": "CreateParameterGroupValue",
	}
	errorCodes := []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1 // a counter used to cover retry scenario; the same below
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "CreateParameterGroup" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if retryIndex >= len(errorCodes)-1 {
						successResponseMock(ReadMockResponseDiff)
						return CreateMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAlicloudRdsParameterGroupCreate(dInit, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_rds_parameter_group"].Schema).Data(dInit.State(), nil)
			for key, value := range attributes {
				dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dInit.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// Update
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewRdsClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAlicloudRdsParameterGroupUpdate(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	// ModifyParameterGroup
	attributesDiff := map[string]interface{}{
		"param_detail": []map[string]interface{}{
			{
				"param_name":  "ModifyParameterGroupValue",
				"param_value": "ModifyParameterGroupValue",
			},
		},
		"parameter_group_desc": "ModifyParameterGroupValue",
		"parameter_group_name": "ModifyParameterGroupValue",
	}
	diff, err := newInstanceDiff("alicloud_rds_parameter_group", attributes, attributesDiff, dInit.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_rds_parameter_group"].Schema).Data(dInit.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeParameterGroup Response
		"ParamGroup": map[string]interface{}{
			"ParameterGroup": []interface{}{
				map[string]interface{}{
					"ParamDetail": map[string]interface{}{
						"ParameterDetail": []interface{}{
							map[string]interface{}{
								"ParamName":  "ModifyParameterGroupValue",
								"ParamValue": "ModifyParameterGroupValue",
							},
						},
					},
					"ParameterGroupName": "ModifyParameterGroupValue",
					"ParameterGroupDesc": "ModifyParameterGroupValue",
				},
			},
		},
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "ModifyParameterGroup" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if retryIndex >= len(errorCodes)-1 {
						return successResponseMock(ReadMockResponseDiff)
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAlicloudRdsParameterGroupUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_rds_parameter_group"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// Read
	errorCodes = []string{"NonRetryableError", "Throttling", "nil", "{}"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DescribeParameterGroup" {
				switch errorCode {
				case "{}":
					return notFoundResponseMock(errorCode)
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if errorCodes[retryIndex] == "nil" {
						return ReadMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAlicloudRdsParameterGroupRead(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "{}":
			assert.Nil(t, err)
		}
	}

	relatedCustinsInfo, ok := dExisted.Get("related_custins_info").([]interface{})
	assert.True(t, ok)
	if ok && len(relatedCustinsInfo) == 1 {
		entry, _ := relatedCustinsInfo[0].(map[string]interface{})
		assert.Equal(t, "CreateParameterGroupValue", entry["db_instance_name"])
		assert.Equal(t, "CreateParameterGroupValue", entry["applied_time"])
	} else {
		assert.Equal(t, 1, len(relatedCustinsInfo))
	}

	// Delete
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewRdsClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAlicloudRdsParameterGroupDelete(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DeleteParameterGroup" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if errorCodes[retryIndex] == "nil" {
						ReadMockResponse = map[string]interface{}{}
						return ReadMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAlicloudRdsParameterGroupDelete(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "nil":
			assert.Nil(t, err)
		}
	}

}

func TestAccAliCloudRdsParameterGroupPostgreSQL(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_parameter_group.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsParameterGroupMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsParameterGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf_testAccAlicloudRdsParameterGroup%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsParameterGroupBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		ProviderFactories: testAccProviderFactory,
		CheckDestroy:      rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"engine":         "PostgreSQL",
					"engine_version": "11.0",
					"param_detail": []map[string]interface{}{
						{
							"param_name":  "enable_sort",
							"param_value": "off",
						},
						{
							"param_name":  "geqo_seed",
							"param_value": "0.1",
						},
					},
					"parameter_group_name": "${var.name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"engine":                 "PostgreSQL",
						"engine_version":         "11.0",
						"param_detail.#":         "2",
						"parameter_group_name":   name,
						"related_custins_info.#": "0",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"param_detail": []map[string]interface{}{
						{
							"param_name":  "enable_sort",
							"param_value": "on",
						},
						{
							"param_name":  "geqo_seed",
							"param_value": "0.2",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"param_detail.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"parameter_group_desc": "update_test",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"parameter_group_desc": "update_test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"parameter_group_name": name + "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"parameter_group_name": name + "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"param_detail": []map[string]interface{}{
						{
							"param_name":  "enable_sort",
							"param_value": "off",
						},
						{
							"param_name":  "geqo_seed",
							"param_value": "0.1",
						},
					},
					"parameter_group_desc": "test",
					"parameter_group_name": "${var.name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"param_detail.#":       "2",
						"parameter_group_desc": "test",
						"parameter_group_name": name,
					}),
				),
			},
		},
	})
}

func TestAccAliCloudRdsParameterGroupResourceGroupAndModifyMode(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_parameter_group.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsParameterGroupMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsParameterGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf_testAccRdsPg%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsParameterGroupResourceGroupDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		ProviderFactories: testAccProviderFactory,
		CheckDestroy:      rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"engine":         "mysql",
					"engine_version": "5.7",
					"param_detail": []map[string]interface{}{
						{
							"param_name":  "back_log",
							"param_value": "3000",
						},
						{
							"param_name":  "wait_timeout",
							"param_value": "86400",
						},
					},
					"parameter_group_name": "${var.name}",
					"resource_group_id":    "${alicloud_resource_manager_resource_group.first.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id":      CHECKSET,
						"related_custins_info.#": "0",
					}),
					testAccCheckRdsParameterGroupInResourceGroup(resourceId, "alicloud_resource_manager_resource_group.first"),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${alicloud_resource_manager_resource_group.second.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
					testAccCheckRdsParameterGroupInResourceGroup(resourceId, "alicloud_resource_manager_resource_group.second"),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"modify_mode": "Individual",
					"param_detail": []map[string]interface{}{
						{
							"param_name":  "back_log",
							"param_value": "4000",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"modify_mode":    "Individual",
						"param_detail.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"modify_mode": "Collectivity",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"modify_mode":    "Collectivity",
						"param_detail.#": "1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"resource_group_id", "modify_mode"},
			},
		},
	})
}

func AlicloudRdsParameterGroupResourceGroupDependence0(name string) string {
	return fmt.Sprintf(`
variable "name" {
  default = "%s"
}

resource "alicloud_resource_manager_resource_group" "first" {
  resource_group_name = "${replace(var.name, "_", "-")}-1"
  display_name        = var.name
}

resource "alicloud_resource_manager_resource_group" "second" {
  resource_group_name = "${replace(var.name, "_", "-")}-2"
  display_name        = var.name
}
`, name)
}

func testAccCheckRdsParameterGroupInResourceGroup(resourceId, resourceGroupRef string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceId]
		if !ok {
			return fmt.Errorf("resource %s not found in state", resourceId)
		}
		resourceGroup, ok := s.RootModule().Resources[resourceGroupRef]
		if !ok {
			return fmt.Errorf("resource group %s not found in state", resourceGroupRef)
		}
		client := testAccProvider.Meta().(*connectivity.AliyunClient)
		response, err := client.RpcPost("Rds", "2014-08-15", "DescribeParameterGroups", map[string]interface{}{
			"RegionId":        client.RegionId,
			"ResourceGroupId": resourceGroup.Primary.ID,
		}, nil, true)
		if err != nil {
			return WrapError(err)
		}
		found := false
		if groups, ok := response["ParameterGroups"].(map[string]interface{}); ok {
			if items, ok := groups["ParameterGroup"].([]interface{}); ok {
				for _, item := range items {
					group, ok := item.(map[string]interface{})
					if !ok {
						continue
					}
					if fmt.Sprint(group["ParameterGroupId"]) == rs.Primary.ID {
						found = true
						break
					}
				}
			}
		}
		if !found {
			return fmt.Errorf("parameter group %s not returned by DescribeParameterGroups with ResourceGroupId %s", rs.Primary.ID, resourceGroup.Primary.ID)
		}
		return nil
	}
}

func TestAccAliCloudRdsParameterGroupRelatedCustinsInfo(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_parameter_group.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsParameterGroupMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsParameterGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf_testAccRdsPg%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsParameterGroupRelatedCustinsInfoDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		ProviderFactories: testAccProviderFactory,
		CheckDestroy:      rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"engine":         "PostgreSQL",
					"engine_version": "15.0",
					"param_detail": []map[string]interface{}{
						{
							"param_name":  "enable_sort",
							"param_value": "off",
						},
						{
							"param_name":  "geqo_seed",
							"param_value": "0.1",
						},
					},
					"parameter_group_name": "${var.name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"engine":                 "PostgreSQL",
						"engine_version":         "15.0",
						"param_detail.#":         "2",
						"parameter_group_name":   name,
						"related_custins_info.#": "1",
						"related_custins_info.0.db_instance_name": CHECKSET,
						"related_custins_info.0.applied_time":     CHECKSET,
					}),
					resource.TestCheckResourceAttrPair(resourceId, "related_custins_info.0.db_instance_name", "alicloud_db_instance.default", "id"),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func AlicloudRdsParameterGroupRelatedCustinsInfoDependence0(name string) string {
	return fmt.Sprintf(`
variable "name" {
  default = "%s"
}

data "alicloud_db_zones" "default" {
  engine                   = "PostgreSQL"
  engine_version           = "15.0"
  instance_charge_type     = "PostPaid"
  category                 = "HighAvailability"
  db_instance_storage_type = "cloud_essd"
}

data "alicloud_db_instance_classes" "default" {
  zone_id                  = data.alicloud_db_zones.default.zones.0.id
  engine                   = "PostgreSQL"
  engine_version           = "15.0"
  category                 = "HighAvailability"
  db_instance_storage_type = "cloud_essd"
  instance_charge_type     = "PostPaid"
}

data "alicloud_vpcs" "default" {
  name_regex = "^default-NODELETING$"
}

data "alicloud_vswitches" "default" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_db_zones.default.zones.0.id
}

resource "alicloud_vswitch" "this" {
  count        = length(data.alicloud_vswitches.default.ids) > 0 ? 0 : 1
  vswitch_name = var.name
  vpc_id       = data.alicloud_vpcs.default.ids.0
  zone_id      = data.alicloud_db_zones.default.ids.0
  cidr_block   = cidrsubnet(data.alicloud_vpcs.default.vpcs.0.cidr_block, 8, 4)
}

locals {
  vswitch_id = length(data.alicloud_vswitches.default.ids) > 0 ? data.alicloud_vswitches.default.ids.0 : concat(alicloud_vswitch.this.*.id, [""])[0]
}

resource "alicloud_db_instance" "default" {
  engine                   = "PostgreSQL"
  engine_version           = "15.0"
  instance_type            = data.alicloud_db_instance_classes.default.instance_classes.0.instance_class
  instance_storage         = data.alicloud_db_instance_classes.default.instance_classes.0.storage_range.0.min
  db_instance_storage_type = "cloud_essd"
  zone_id                  = data.alicloud_db_zones.default.zones.0.id
  vswitch_id               = local.vswitch_id
  instance_charge_type     = "Postpaid"
  instance_name            = "${var.name}db"
  db_param_group_id        = alicloud_rds_parameter_group.default.id
}
`, name)
}

func TestUnitAlicloudRdsParameterGroupRequestParameters(t *testing.T) {
	p := Provider().ResourcesMap
	region := os.Getenv("ALICLOUD_REGION")
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		t.Skipf("Skipping the test case with err: %s", err)
		t.Skipped()
	}
	rawClient = rawClient.(*connectivity.AliyunClient)

	describeResponse := func(id string) map[string]interface{} {
		return map[string]interface{}{
			"ParamGroup": map[string]interface{}{
				"ParameterGroup": []interface{}{
					map[string]interface{}{
						"ParameterGroupId": id,
						"Engine":           "mysql",
						"EngineVersion":    "5.7",
						"ParamDetail": map[string]interface{}{
							"ParameterDetail": []interface{}{
								map[string]interface{}{
									"ParamName":  "back_log",
									"ParamValue": "3000",
								},
							},
						},
						"ParameterGroupName": "tf-testacc-parameter-group",
					},
				},
			},
		}
	}

	var captured map[string]interface{}
	patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, body map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
		switch *action {
		case "CreateParameterGroup":
			captured = body
			return map[string]interface{}{"ParameterGroupId": "rpg-test"}, nil
		case "ModifyParameterGroup":
			captured = body
			return map[string]interface{}{"ParameterGroupId": "rpg-test"}, nil
		}
		return describeResponse("rpg-test"), nil
	})
	defer patches.Reset()

	newAttributes := func() map[string]interface{} {
		return map[string]interface{}{
			"engine":         "mysql",
			"engine_version": "5.7",
			"param_detail": []map[string]interface{}{
				{
					"param_name":  "back_log",
					"param_value": "3000",
				},
			},
			"parameter_group_name": "tf-testacc-parameter-group",
		}
	}

	dCreate, _ := schema.InternalMap(p["alicloud_rds_parameter_group"].Schema).Data(nil, nil)
	dCreate.MarkNewResource()
	for key, value := range newAttributes() {
		assert.Nil(t, dCreate.Set(key, value))
	}
	assert.Nil(t, dCreate.Set("resource_group_id", "rg-test"))
	captured = nil
	assert.Nil(t, resourceAlicloudRdsParameterGroupCreate(dCreate, rawClient))
	assert.Equal(t, "rg-test", captured["ResourceGroupId"])
	if _, ok := captured["ModifyMode"]; ok {
		t.Fatal("CreateParameterGroup must not send ModifyMode")
	}

	attributes := newAttributes()
	diff, err := newInstanceDiff("alicloud_rds_parameter_group", attributes, map[string]interface{}{
		"resource_group_id": "rg-test2",
		"modify_mode":       "Individual",
	}, dCreate.State())
	assert.Nil(t, err)
	dUpdate, _ := schema.InternalMap(p["alicloud_rds_parameter_group"].Schema).Data(dCreate.State(), diff)
	captured = nil
	assert.Nil(t, resourceAlicloudRdsParameterGroupUpdate(dUpdate, rawClient))
	assert.Equal(t, "rg-test2", captured["ResourceGroupId"])
	assert.Equal(t, "Individual", captured["ModifyMode"])

	attributes = newAttributes()
	attributes["resource_group_id"] = "rg-test2"
	attributes["modify_mode"] = "Individual"
	diff, err = newInstanceDiff("alicloud_rds_parameter_group", attributes, map[string]interface{}{
		"parameter_group_desc": "updated",
	}, dCreate.State())
	assert.Nil(t, err)
	dUpdate, _ = schema.InternalMap(p["alicloud_rds_parameter_group"].Schema).Data(dUpdate.State(), diff)
	captured = nil
	assert.Nil(t, resourceAlicloudRdsParameterGroupUpdate(dUpdate, rawClient))
	if _, ok := captured["ModifyMode"]; ok {
		t.Fatal("ModifyParameterGroup must not send an unchanged ModifyMode")
	}
	if _, ok := captured["ResourceGroupId"]; ok {
		t.Fatal("ModifyParameterGroup must not send an unchanged ResourceGroupId")
	}
}
