package cloudfirewall_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/PaesslerAG/jsonpath"
	credentials "github.com/aliyun/credentials-go/credentials"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/acctest"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/errs"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/helper"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/logs"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/service/cloudfirewall"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdkterraform "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tfacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// collection-order: bypass asset_region_resource_types -- pre-existing list unchanged by this PR; regression changes cover address_list unknown handling only

func TestAccAliCloudCloudFirewallAddressBook_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	checkoutSupportedRegions(t, true, connectivity.CloudFirewallSupportRegions)
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return cloudfirewall.NewService(acctest.Provider.Meta().(*connectivity.AliyunClient))
	}, "DescribeAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudfirewalladdressbook%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName:            resourceId,
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		CheckDestroy:             rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":   name,
					"group_type":   "ip",
					"description":  name,
					"address_list": []string{"10.21.0.0/16", "10.22.0.0/16", "10.168.0.0/16"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     name,
						"group_type":     "ip",
						"description":    name,
						"address_list.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"lang": "en",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"lang": "en",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"address_list": []string{"10.21.0.0/16"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"address_list.#": "1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

func TestAccAliCloudCloudFirewallAddressBook_basic1(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	checkoutSupportedRegions(t, true, connectivity.CloudFirewallSupportRegions)
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return cloudfirewall.NewService(acctest.Provider.Meta().(*connectivity.AliyunClient))
	}, "DescribeAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudfirewalladdressbook%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName:            resourceId,
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		CheckDestroy:             rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":   name,
					"group_type":   "ipv6",
					"description":  name,
					"address_list": []string{"::1/128", "::2/128", "::3/128"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     name,
						"group_type":     "ipv6",
						"description":    name,
						"address_list.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"lang": "en",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"lang": "en",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"address_list": []string{"::1/128"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"address_list.#": "1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

func TestAccAliCloudCloudFirewallAddressBook_basic2(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	checkoutSupportedRegions(t, true, connectivity.CloudFirewallSupportRegions)
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return cloudfirewall.NewService(acctest.Provider.Meta().(*connectivity.AliyunClient))
	}, "DescribeAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudfirewalladdressbook%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName:            resourceId,
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		CheckDestroy:             rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":   name,
					"group_type":   "domain",
					"description":  name,
					"address_list": []string{"alibaba.com", "aliyun.com", "alicloud.com"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     name,
						"group_type":     "domain",
						"description":    name,
						"address_list.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"lang": "en",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"lang": "en",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"address_list": []string{"alibaba.com"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"address_list.#": "1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

func TestAccAliCloudCloudFirewallAddressBook_basic3(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	checkoutSupportedRegions(t, true, connectivity.CloudFirewallSupportRegions)
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return cloudfirewall.NewService(acctest.Provider.Meta().(*connectivity.AliyunClient))
	}, "DescribeAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudfirewalladdressbook%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName:            resourceId,
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		CheckDestroy:             rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":   name,
					"group_type":   "port",
					"description":  name,
					"address_list": []string{"1/1", "22/22", "88/88"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     name,
						"group_type":     "port",
						"description":    name,
						"address_list.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"lang": "en",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"lang": "en",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"address_list": []string{"1/1"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"address_list.#": "1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

func TestAccAliCloudCloudFirewallAddressBook_basic4(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	checkoutSupportedRegions(t, true, connectivity.CloudFirewallSupportRegions)
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return cloudfirewall.NewService(acctest.Provider.Meta().(*connectivity.AliyunClient))
	}, "DescribeAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudfirewalladdressbook%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName:            resourceId,
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		CheckDestroy:             rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":  name,
					"group_type":  "tag",
					"description": name,
					"ecs_tags": []map[string]interface{}{
						{
							"tag_key":   "created",
							"tag_value": "tfTestAcc0",
						},
						{
							"tag_key":   "for",
							"tag_value": "tfTestAcc1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":  name,
						"group_type":  "tag",
						"description": name,
						"ecs_tags.#":  "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_add_tag_ecs": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_add_tag_ecs": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tag_relation": "or",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tag_relation": "or",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"lang": "en",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"lang": "en",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ecs_tags": []map[string]interface{}{
						{
							"tag_key":   "created",
							"tag_value": "tfTestAcc0",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ecs_tags.#": "1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

func TestAccAliCloudCloudFirewallAddressBook_basic5(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	checkoutSupportedRegions(t, true, connectivity.CloudFirewallSupportRegions)
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return cloudfirewall.NewService(acctest.Provider.Meta().(*connectivity.AliyunClient))
	}, "DescribeAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudfirewalladdressbook%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName:            resourceId,
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		CheckDestroy:             rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":       name,
					"group_type":       "tag",
					"description":      name,
					"auto_add_tag_ecs": "1",
					"tag_relation":     "or",
					"lang":             "en",
					"ecs_tags": []map[string]interface{}{
						{
							"tag_key":   "created",
							"tag_value": "tfTestAcc0",
						},
						{
							"tag_key":   "for",
							"tag_value": "tfTestAcc1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":       name,
						"group_type":       "tag",
						"description":      name,
						"auto_add_tag_ecs": "1",
						"tag_relation":     "or",
						"lang":             "en",
						"ecs_tags.#":       "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_add_tag_ecs": "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_add_tag_ecs": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tag_relation": "and",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tag_relation": "and",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"lang": "zh",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"lang": "zh",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ecs_tags": []map[string]interface{}{
						{
							"tag_key":   "created",
							"tag_value": "tfTestAcc0",
						},
						{
							"tag_key":   "for",
							"tag_value": "tfTestAcc1",
						},
						{
							"tag_key":   "by",
							"tag_value": "tfTestAcc2",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ecs_tags.#": "3",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

func TestAccAliCloudCloudFirewallAddressBook_basic6(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	checkoutSupportedRegions(t, true, connectivity.CloudFirewallSupportRegions)
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return cloudfirewall.NewService(acctest.Provider.Meta().(*connectivity.AliyunClient))
	}, "DescribeAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudfirewalladdressbook%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName:            resourceId,
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		CheckDestroy:             rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":  name,
					"group_type":  "asset",
					"description": name,
					"asset_region_resource_types": []map[string]interface{}{
						{
							"asset_region_id": "all",
							"resource_type": []map[string]interface{}{
								{
									"ipv4": []map[string]interface{}{
										{
											"eip":                     "false",
											"ecs_eip":                 "false",
											"ecs_public_ip":           "false",
											"slb_eip":                 "false",
											"slb_public_ip":           "false",
											"nlb_eip":                 "false",
											"alb_eip":                 "false",
											"nat_eip":                 "false",
											"nat_public_ip":           "false",
											"eni_eip":                 "false",
											"ga_eip":                  "true",
											"api_gateway_eip":         "false",
											"ai_gateway_eip":          "false",
											"bastion_host_ip":         "true",
											"bastion_host_ingress_ip": "true",
											"bastion_host_egress_ip":  "true",
											"havip":                   "true",
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":                    name,
						"group_type":                    "asset",
						"description":                   name,
						"asset_region_resource_types.#": "1",
						"asset_region_resource_types.0.asset_region_id":                        "all",
						"asset_region_resource_types.0.resource_type.#":                        "1",
						"asset_region_resource_types.0.resource_type.0.ipv4.#":                 "1",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.ga_eip":          "true",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.bastion_host_ip": "true",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.havip":           "true",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.eip":             "false",
						"address_list_count": CHECKSET,
						"reference_count":    CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": name + "update",
					"asset_region_resource_types": []map[string]interface{}{
						{
							"asset_region_id": "all",
							"resource_type": []map[string]interface{}{
								{
									"ipv4": []map[string]interface{}{
										{
											"eip":                     "true",
											"ecs_eip":                 "true",
											"ecs_public_ip":           "true",
											"slb_eip":                 "true",
											"slb_public_ip":           "true",
											"nlb_eip":                 "true",
											"alb_eip":                 "true",
											"nat_eip":                 "true",
											"nat_public_ip":           "true",
											"eni_eip":                 "true",
											"ga_eip":                  "false",
											"api_gateway_eip":         "true",
											"ai_gateway_eip":          "true",
											"bastion_host_ip":         "false",
											"bastion_host_ingress_ip": "false",
											"bastion_host_egress_ip":  "false",
											"havip":                   "false",
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name": name + "update",
						"asset_region_resource_types.0.asset_region_id":                        "all",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.eip":             "true",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.ecs_eip":         "true",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.ecs_public_ip":   "true",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.slb_eip":         "true",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.slb_public_ip":   "true",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.nlb_eip":         "true",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.alb_eip":         "true",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.nat_eip":         "true",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.nat_public_ip":   "true",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.eni_eip":         "true",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.ga_eip":          "false",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.api_gateway_eip": "true",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.ai_gateway_eip":  "true",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.bastion_host_ip": "false",
						"asset_region_resource_types.0.resource_type.0.ipv4.0.havip":           "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"asset_member_uids": []string{"${data.alicloud_account.default.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"asset_member_uids.#": "1",
						"asset_member_uids.0": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang", "asset_region_resource_types.asset_region_id"},
			},
		},
	})
}

func TestAccAliCloudCloudFirewallAddressBook_basic7(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	checkoutSupportedRegions(t, true, connectivity.CloudFirewallSupportRegions)
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return cloudfirewall.NewService(acctest.Provider.Meta().(*connectivity.AliyunClient))
	}, "DescribeAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudfirewalladdressbook%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName:            resourceId,
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		CheckDestroy:             rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":  name,
					"group_type":  "assetIpv6",
					"description": name,
					"asset_region_resource_types": []map[string]interface{}{
						{
							"asset_region_id": "all",
							"resource_type": []map[string]interface{}{
								{
									"ipv6": []map[string]interface{}{
										{
											"ecs_ipv6":          "false",
											"slb_ipv6":          "false",
											"nlb_ipv6":          "false",
											"alb_ipv6":          "false",
											"eni_eipv6":         "false",
											"ga_eipv6":          "true",
											"api_gateway_eipv6": "false",
											"ai_gateway_eipv6":  "false",
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":                    name,
						"group_type":                    "assetIpv6",
						"description":                   name,
						"asset_region_resource_types.#": "1",
						"asset_region_resource_types.0.asset_region_id":                 "all",
						"asset_region_resource_types.0.resource_type.#":                 "1",
						"asset_region_resource_types.0.resource_type.0.ipv6.#":          "1",
						"asset_region_resource_types.0.resource_type.0.ipv6.0.ga_eipv6": "true",
						"asset_region_resource_types.0.resource_type.0.ipv6.0.ecs_ipv6": "false",
						"address_list_count": CHECKSET,
						"reference_count":    CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"asset_region_resource_types": []map[string]interface{}{
						{
							"asset_region_id": "all",
							"resource_type": []map[string]interface{}{
								{
									"ipv6": []map[string]interface{}{
										{
											"ecs_ipv6":          "true",
											"slb_ipv6":          "true",
											"nlb_ipv6":          "true",
											"alb_ipv6":          "true",
											"eni_eipv6":         "true",
											"ga_eipv6":          "false",
											"api_gateway_eipv6": "true",
											"ai_gateway_eipv6":  "true",
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"asset_region_resource_types.0.resource_type.0.ipv6.0.ecs_ipv6":          "true",
						"asset_region_resource_types.0.resource_type.0.ipv6.0.slb_ipv6":          "true",
						"asset_region_resource_types.0.resource_type.0.ipv6.0.nlb_ipv6":          "true",
						"asset_region_resource_types.0.resource_type.0.ipv6.0.alb_ipv6":          "true",
						"asset_region_resource_types.0.resource_type.0.ipv6.0.eni_eipv6":         "true",
						"asset_region_resource_types.0.resource_type.0.ipv6.0.ga_eipv6":          "false",
						"asset_region_resource_types.0.resource_type.0.ipv6.0.api_gateway_eipv6": "true",
						"asset_region_resource_types.0.resource_type.0.ipv6.0.ai_gateway_eipv6":  "true",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AliCloudCloudFirewallAddressBookMap0 = map[string]string{}

func AliCloudCloudFirewallAddressBookBasicDependence0(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_account" "default" {
	}
`, name)
}

// Exercise the real Read and SDK diff together: API response order must not
// create changes, while membership, multiplicity and complete values still do.
func TestUnitCloudFirewallAddressBookRefreshDiff(t *testing.T) {
	a, b, c := "192.0.2.1/32", "198.51.100.1/32", "203.0.113.1/32"
	for _, tc := range []struct {
		name             string
		previous, remote []interface{}
		changed          bool
	}{
		{"same order", []interface{}{a, b, c}, []interface{}{a, b, c}, false},
		{"reverse", []interface{}{a, b, c}, []interface{}{c, b, a}, false},
		{"last to first", []interface{}{a, b, c}, []interface{}{c, a, b}, false},
		{"duplicate reorder", []interface{}{a, a, b}, []interface{}{b, a, a}, false},
		{"added", []interface{}{a, b}, []interface{}{c, b, a}, true},
		{"removed", []interface{}{a, b, c}, []interface{}{c, a}, true},
		{"replaced", []interface{}{a, b}, []interface{}{a, c}, true},
		{"duplicate count", []interface{}{a, a, b}, []interface{}{a, b, b}, true},
		{"description", []interface{}{a + " web", b}, []interface{}{b, a + " database"}, true},
		{"whitespace", []interface{}{a + " web", b}, []interface{}{b, a + "  web"}, true},
		{"case", []interface{}{"example.com Web", b}, []interface{}{b, "example.com web"}, true},
		{"empty member", []interface{}{a, ""}, []interface{}{a, b}, true},
		{"cleared", []interface{}{a, b}, []interface{}{}, true},
		{"empty", []interface{}{}, []interface{}{}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := cloudfirewall.ResourceAddressBook()
			config := map[string]interface{}{"group_name": "test", "group_type": "ip", "description": "test", "address_list": tc.previous}
			d := schema.TestResourceDataRaw(t, r.Schema, config)
			d.SetId("address-book")
			client := cloudFirewallAddressBookTestClient(t, tc.remote)
			refreshed, diags := r.RefreshWithoutUpgrade(context.Background(), d.State(), client)
			if diags.HasError() {
				t.Fatalf("%v", diags)
			}
			diff, err := r.Diff(context.Background(), refreshed, sdkterraform.NewResourceConfigRaw(config), nil)
			if err != nil {
				t.Fatal(err)
			}
			changed := diff != nil && !diff.Empty()
			if changed != tc.changed {
				t.Fatalf("changed=%v want=%v diff=%#v", changed, tc.changed, diff)
			}
		})
	}
}

func cloudFirewallAddressBookTestClient(t *testing.T, addresses []interface{}, inspect ...func(*http.Request)) *connectivity.AliyunClient {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		for _, check := range inspect {
			check(req)
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"GroupUuid": "address-book", "Acls": []interface{}{map[string]interface{}{
			"GroupUuid": "address-book", "GroupName": "test", "GroupType": "ip", "Description": "test", "AddressList": addresses, "TagList": []interface{}{},
		}}}); err != nil {
			t.Error(err)
		}
	}))
	t.Cleanup(server.Close)
	credential, err := credentials.NewCredential(new(credentials.Config).SetType("access_key").SetAccessKeyId("test-key").SetAccessKeySecret("test-secret"))
	if err != nil {
		t.Fatal(err)
	}
	endpoint := strings.TrimPrefix(server.URL, "http://")
	t.Setenv("NO_PROXY", endpoint)
	endpoints := new(sync.Map)
	config := &connectivity.Config{AccessKey: "test-key", SecretKey: "test-secret", Credential: credential, RegionId: "cn-hangzhou", AccountType: "test", Protocol: "http", Endpoints: endpoints, SignVersion: new(sync.Map), SkipRegionValidation: true}
	client, err := config.Client()
	if err != nil {
		t.Fatal(err)
	}
	endpoints.Store("cloudfw", endpoint)
	return client
}

func TestUnitCloudFirewallAddressBookConfiguredDiff(t *testing.T) {
	a, b := "192.0.2.1/32", "198.51.100.1/32"
	for _, tc := range []struct {
		name        string
		addresses   interface{}
		omitted     bool
		description string
		changed     bool
	}{
		{"same", []interface{}{a, b}, false, "test", false},
		{"configured reorder", []interface{}{b, a}, false, "test", true},
		{"computed omitted", nil, true, "test", false},
		// The SDK plans an unknown count for an entirely unknown Optional+Computed list.
		{"whole list unknown", "74D93920-ED26-11E3-AC10-0800200C9A66", false, "test", true},
		{"unknown element", []interface{}{a, "74D93920-ED26-11E3-AC10-0800200C9A66"}, false, "test", true},
		{"other change", []interface{}{a, b}, false, "updated", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := cloudfirewall.ResourceAddressBook()
			d := schema.TestResourceDataRaw(t, r.Schema, map[string]interface{}{"group_name": "test", "group_type": "ip", "description": "test", "address_list": []interface{}{a, b}})
			d.SetId("address-book")
			config := map[string]interface{}{"group_name": "test", "group_type": "ip", "description": tc.description}
			if !tc.omitted {
				config["address_list"] = tc.addresses
			}
			diff, err := r.Diff(context.Background(), d.State(), sdkterraform.NewResourceConfigRaw(config), nil)
			if err != nil {
				t.Fatal(err)
			}
			if changed := diff != nil && !diff.Empty(); changed != tc.changed {
				t.Fatalf("changed=%v want=%v diff=%#v", changed, tc.changed, diff)
			}
			if tc.name == "whole list unknown" {
				count := diff.Attributes["address_list.#"]
				if count == nil || !count.NewComputed {
					t.Fatal("unknown list count must remain computed in the plan")
				}
				if len(diff.Attributes) != 1 {
					t.Fatalf("unknown list must not produce concrete element changes: %#v", diff.Attributes)
				}
				client := cloudFirewallAddressBookTestClient(t, []interface{}{a, b}, func(req *http.Request) {
					if err := req.ParseForm(); err != nil {
						t.Error(err)
					}
					if req.Form.Get("Action") == "ModifyAddressBook" && req.Form.Has("AddressList") {
						t.Errorf("unknown list must not be sent to the API: %q", req.Form.Get("AddressList"))
					}
				})
				applied, diags := r.Apply(context.Background(), d.State(), diff, client)
				if diags.HasError() {
					t.Fatalf("apply unknown list: %v", diags)
				}
				if applied.Attributes["address_list.#"] != "2" || applied.Attributes["address_list.0"] != a || applied.Attributes["address_list.1"] != b {
					t.Fatalf("apply must preserve remote addresses: %#v", applied.Attributes)
				}
			}
			if tc.name == "other change" && diff.Attributes["description"] == nil {
				t.Fatal("description diff lost")
			}
		})
	}
}

func TestUnitCloudFirewallAddressBookApplyOrder(t *testing.T) {
	r := cloudfirewall.ResourceAddressBook()
	a, b := "192.0.2.1/32", "198.51.100.1/32"
	config := map[string]interface{}{"group_name": "test", "group_type": "ip", "description": "test", "address_list": []interface{}{a, b}}
	d := schema.TestResourceDataRaw(t, r.Schema, config)
	d.SetId("address-book")
	client := cloudFirewallAddressBookTestClient(t, []interface{}{a, b})
	for _, addresses := range [][]interface{}{{b, a}, {a, b}} {
		config["address_list"] = addresses
		diff, err := r.Diff(context.Background(), d.State(), sdkterraform.NewResourceConfigRaw(config), nil)
		if err != nil {
			t.Fatal(err)
		}
		if diff == nil || diff.Empty() {
			t.Fatal("configured order change must have a plan")
		}
		applied, diags := r.Apply(context.Background(), d.State(), diff, client)
		if diags.HasError() {
			t.Fatalf("%v", diags)
		}
		refreshed, diags := r.RefreshWithoutUpgrade(context.Background(), applied, client)
		if diags.HasError() {
			t.Fatalf("%v", diags)
		}
		diff, err = r.Diff(context.Background(), refreshed, sdkterraform.NewResourceConfigRaw(config), nil)
		if err != nil {
			t.Fatal(err)
		}
		if diff != nil && !diff.Empty() {
			t.Fatalf("order did not converge after apply and refresh: %#v", diff)
		}
		d = r.Data(refreshed)
	}
}

func TestAccAliCloudCloudFirewallAddressBook_addressListOrder(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	name := fmt.Sprintf("tf-testacc-address-order-%d", tfacctest.RandIntRange(10000, 99999))
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return cloudfirewall.NewService(acctest.Provider.Meta().(*connectivity.AliyunClient))
	}, "DescribeAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheckWithRegions(t, true, connectivity.CloudFirewallSupportRegions) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		CheckDestroy:             rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{Config: testAccConfig(map[string]interface{}{"group_name": name, "group_type": "ip", "description": name, "address_list": []string{"192.0.2.1/32", "198.51.100.1/32", "203.0.113.1/32"}})},
			{Config: testAccConfig(map[string]interface{}{"address_list": []string{"203.0.113.1/32", "192.0.2.1/32", "198.51.100.1/32"}}), PlanOnly: true, ExpectNonEmptyPlan: true},
			{Config: testAccConfig(map[string]interface{}{"address_list": []string{"203.0.113.1/32", "192.0.2.1/32", "198.51.100.1/32"}}), Check: resource.TestCheckResourceAttr(resourceId, "address_list.0", "203.0.113.1/32")},
			{Config: testAccConfig(map[string]interface{}{"address_list": []string{"192.0.2.1/32", "198.51.100.1/32", "203.0.113.1/32"}}), PlanOnly: true, ExpectNonEmptyPlan: true},
			{Config: testAccConfig(map[string]interface{}{"address_list": []string{"192.0.2.1/32", "198.51.100.1/32", "203.0.113.1/32"}}), Check: resource.TestCheckResourceAttr(resourceId, "address_list.0", "192.0.2.1/32")},
			{Config: testAccConfig(map[string]interface{}{"address_list": []string{"192.0.2.1/32", "198.51.100.1/32", "203.0.113.1/32", "192.0.2.2/32"}}), Check: resource.TestCheckResourceAttr(resourceId, "address_list.#", "4")},
			{Config: testAccConfig(map[string]interface{}{"address_list": []string{"192.0.2.1/32", "198.51.100.1/32"}}), Check: resource.TestCheckResourceAttr(resourceId, "address_list.#", "2")},
			{Config: testAccConfig(map[string]interface{}{"address_list": []string{"192.0.2.1/32", "203.0.113.1/32"}}), Check: resource.TestCheckResourceAttr(resourceId, "address_list.1", "203.0.113.1/32")},
			// Imported lists have no prior configured order. Verify membership below;
			// and require the following configured apply to converge after refresh.
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"address_list"},
				ImportStateCheck: func(states []*terraform.InstanceState) error {
					if len(states) != 1 || states[0].Attributes["address_list.#"] != "2" {
						return fmt.Errorf("expected two imported addresses")
					}
					values := map[string]bool{states[0].Attributes["address_list.0"]: true, states[0].Attributes["address_list.1"]: true}
					if !values["192.0.2.1/32"] || !values["203.0.113.1/32"] {
						return fmt.Errorf("imported addresses do not match configuration")
					}
					return nil
				}},
			{Config: testAccConfig(map[string]interface{}{})},
		},
	})
}

// Create an isolated member instead of depending on accounts already managed
// by the acceptance environment. Terraform destroys the address book, removes
// this member, then submits deletion of the newly created resource account.
func TestAccAliCloudCloudFirewallAddressBook_assetMemberUidsOrder(t *testing.T) {
	var v, member map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	memberResourceId := "alicloud_cloud_firewall_instance_member.order"
	name := fmt.Sprintf("tf-testacc-member-order-%d", tfacctest.RandIntRange(10000, 99999))
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return cloudfirewall.NewService(acctest.Provider.Meta().(*connectivity.AliyunClient))
	}, "DescribeAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	memberCheck := resourceCheckInitWithDescribeMethod(memberResourceId, &member, func() interface{} {
		return &cloudFirewallInstanceMemberTestService{client: acctest.Provider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstanceMember")
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, cloudFirewallAddressBookMemberOrderDependence)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheckWithRegions(t, true, connectivity.CloudFirewallSupportRegions) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		CheckDestroy:             resource.ComposeTestCheckFunc(rac.checkResourceDestroy(), memberCheck.checkResourceDestroy()),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":                  name,
					"group_type":                  "asset",
					"description":                 name,
					"asset_member_uids":           []string{"${data.alicloud_account.default.id}", "${alicloud_cloud_firewall_instance_member.order.member_uid}"},
					"asset_region_resource_types": []map[string]interface{}{{"asset_region_id": "all", "resource_type": []map[string]interface{}{{"ipv4": []map[string]interface{}{{"eip": "true"}}}}}},
				}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(memberResourceId, "status", "normal"),
					resource.TestCheckResourceAttrPair(memberResourceId, "member_uid", "alicloud_resource_manager_account.order", "id"),
					resource.TestCheckResourceAttr(resourceId, "asset_member_uids.#", "2"),
					resource.TestCheckResourceAttrPair(resourceId, "asset_member_uids.0", "data.alicloud_account.default", "id"),
					resource.TestCheckResourceAttrPair(resourceId, "asset_member_uids.1", memberResourceId, "member_uid"),
					func(s *terraform.State) error {
						attrs := s.RootModule().Resources[resourceId].Primary.Attributes
						if attrs["asset_member_uids.0"] == attrs["asset_member_uids.1"] {
							return fmt.Errorf("isolated member ordering test requires two distinct account IDs")
						}
						t.Log("isolated member fixture is normal and distinct from the current account")
						return nil
					},
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"asset_member_uids": []string{"${alicloud_cloud_firewall_instance_member.order.member_uid}", "${data.alicloud_account.default.id}"},
				}),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"asset_member_uids": []string{"${alicloud_cloud_firewall_instance_member.order.member_uid}", "${data.alicloud_account.default.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceId, "asset_member_uids.#", "2"),
					resource.TestCheckResourceAttrPair(resourceId, "asset_member_uids.0", memberResourceId, "member_uid"),
					resource.TestCheckResourceAttrPair(resourceId, "asset_member_uids.1", "data.alicloud_account.default", "id"),
				),
			},
		},
	})
}

func cloudFirewallAddressBookMemberOrderDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
  default = "%s"
}

data "alicloud_account" "default" {}

resource "alicloud_resource_manager_account" "order" {
  display_name = var.name
}

resource "alicloud_cloud_firewall_instance_member" "order" {
  member_uid = alicloud_resource_manager_account.order.id
  member_desc = var.name
}
`, name)
}

func TestUnitCloudFirewallAddressBookImportOrder(t *testing.T) {
	r := cloudfirewall.ResourceAddressBook()
	client := cloudFirewallAddressBookTestClient(t, []interface{}{"198.51.100.1/32", "192.0.2.1/32", "192.0.2.1/32"})
	refreshed, diags := r.RefreshWithoutUpgrade(context.Background(), &sdkterraform.InstanceState{ID: "address-book", Attributes: map[string]string{}}, client)
	if diags.HasError() {
		t.Fatalf("%v", diags)
	}
	for key, want := range map[string]string{"address_list.#": "3", "address_list.0": "198.51.100.1/32", "address_list.1": "192.0.2.1/32", "address_list.2": "192.0.2.1/32"} {
		if got := refreshed.Attributes[key]; got != want {
			t.Fatalf("%s=%q want %q", key, got, want)
		}
	}
}

func TestUnitCloudFirewallAddressBookNewOccurrencesOrder(t *testing.T) {
	r := cloudfirewall.ResourceAddressBook()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]interface{}{"group_name": "test", "group_type": "ip", "description": "test", "address_list": []interface{}{"192.0.2.1/32"}})
	d.SetId("address-book")
	client := cloudFirewallAddressBookTestClient(t, []interface{}{"192.0.2.1/32", "198.51.100.1/32", "192.0.2.1/32"})
	refreshed, diags := r.RefreshWithoutUpgrade(context.Background(), d.State(), client)
	if diags.HasError() {
		t.Fatalf("%v", diags)
	}
	for key, want := range map[string]string{"address_list.#": "3", "address_list.0": "192.0.2.1/32", "address_list.1": "198.51.100.1/32", "address_list.2": "192.0.2.1/32"} {
		if got := refreshed.Attributes[key]; got != want {
			t.Fatalf("%s=%q want %q", key, got, want)
		}
	}
}

// Observe the service response directly, then keep configuration unchanged while
// Terraform refreshes and plans. Two fixed orders cover services that preserve
// submitted order as well as services that always normalize to one order.
func TestAccAliCloudCloudFirewallAddressBook_addressListRemoteOrder(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	name := fmt.Sprintf("tf-testacc-remote-order-%d", tfacctest.RandIntRange(10000, 99999))
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return cloudfirewall.NewService(acctest.Provider.Meta().(*connectivity.AliyunClient))
	}, "DescribeAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
	observedDifferentOrder := false
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheckWithRegions(t, true, connectivity.CloudFirewallSupportRegions) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		CheckDestroy:             rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":   name,
					"group_type":   "ip",
					"description":  name,
					"address_list": []string{"192.0.2.1/32", "198.51.100.1/32", "203.0.113.1/32"},
				}),
				Check: testAccCloudFirewallAddressBookObserveRemoteOrder(t, resourceId,
					[]string{"192.0.2.1/32", "198.51.100.1/32", "203.0.113.1/32"}, &observedDifferentOrder),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"address_list": []string{"192.0.2.1/32", "198.51.100.1/32", "203.0.113.1/32"},
				}),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				PreConfig: func() { t.Log("unchanged configuration A empty-plan step completed") },
				Config: testAccConfig(map[string]interface{}{
					"address_list": []string{"203.0.113.1/32", "198.51.100.1/32", "192.0.2.1/32"},
				}),
				Check: testAccCloudFirewallAddressBookObserveRemoteOrder(t, resourceId,
					[]string{"203.0.113.1/32", "198.51.100.1/32", "192.0.2.1/32"}, &observedDifferentOrder),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"address_list": []string{"203.0.113.1/32", "198.51.100.1/32", "192.0.2.1/32"},
				}),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: func(s *terraform.State) error {
					t.Log("unchanged configuration B empty-plan step completed")
					if !observedDifferentOrder {
						return fmt.Errorf("test did not observe remote address order differing from unchanged configuration")
					}
					t.Log("verified remote address order differs while unchanged configuration plans remain empty")
					return nil
				},
			},
		},
	})
}

func testAccCloudFirewallAddressBookObserveRemoteOrder(t *testing.T, resourceId string, configured []string, observed *bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		t.Helper()
		rs, ok := s.RootModule().Resources[resourceId]
		if !ok || rs.Primary == nil || rs.Primary.ID == "" {
			return fmt.Errorf("test address book is missing from state")
		}
		client := acctest.Provider.Meta().(*connectivity.AliyunClient)
		service := cloudfirewall.NewService(client)
		readDifferentOrder := func() (bool, error) {
			object, err := service.DescribeAddressBook(rs.Primary.ID)
			if err != nil {
				return false, err
			}
			addresses, ok := object["AddressList"].([]interface{})
			if !ok || len(addresses) != len(configured) {
				return false, fmt.Errorf("remote address membership changed during ordering test")
			}
			remaining := make(map[string]int, len(configured))
			for _, address := range configured {
				remaining[address]++
			}
			differs := false
			for i, raw := range addresses {
				address, ok := raw.(string)
				if !ok || remaining[address] == 0 {
					return false, fmt.Errorf("remote address membership changed during ordering test")
				}
				remaining[address]--
				differs = differs || address != configured[i]
			}
			return differs, nil
		}
		differs, err := readDifferentOrder()
		if err != nil {
			return err
		}
		if !differs {
			reversed := make([]string, len(configured))
			for i, address := range configured {
				reversed[len(configured)-1-i] = address
			}
			request := map[string]interface{}{
				"GroupUuid":   rs.Primary.ID,
				"GroupName":   rs.Primary.Attributes["group_name"],
				"Description": rs.Primary.Attributes["description"],
				"AddressList": strings.Join(reversed, ","),
			}
			endpoint := ""
			err = resource.Retry(time.Minute, func() *resource.RetryError {
				_, err := client.RpcPostWithEndpoint("Cloudfw", "2017-12-07", "ModifyAddressBook", nil, request, false, endpoint)
				if err == nil {
					return nil
				}
				if errs.IsExpectedErrors(err, []string{"not buy user"}) {
					endpoint = connectivity.CloudFirewallOpenAPIEndpointControlPolicy
					return resource.RetryableError(err)
				}
				if helper.NeedRetry(err) {
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			})
			if err != nil {
				return err
			}
			differs, err = readDifferentOrder()
			if err != nil {
				return err
			}
		}
		t.Logf("remote address order differs from unchanged configuration: %t", differs)
		*observed = *observed || differs
		return nil
	}
}

// The CloudFirewall acceptance tests used the root package helpers before the
// migration. The split package keeps the same region selection: CloudFirewall
// only exists in ap-southeast-1 and cn-hangzhou, so an unsupported ALICLOUD_REGION
// is replaced before any name or provider is built.
var defaultRegionToTest = func() string {
	if v := os.Getenv("ALICLOUD_REGION"); v != "" {
		return v
	}
	return "cn-beijing"
}()

func testAccPreCheck(t *testing.T) {
	acctest.PreCheck(t)
	defaultRegionToTest = os.Getenv("ALICLOUD_REGION")
}

func testAccPreCheckWithRegions(t *testing.T, supported bool, regions []connectivity.Region) {
	acctest.PreCheck(t)
	checkoutSupportedRegions(t, supported, regions)
	defaultRegionToTest = os.Getenv("ALICLOUD_REGION")
}

func checkoutSupportedRegions(t *testing.T, supported bool, regions []connectivity.Region) {
	region := os.Getenv("ALICLOUD_REGION")
	find := false
	backupRegion := string(connectivity.APSouthEast1)
	if region == string(connectivity.APSouthEast1) {
		backupRegion = string(connectivity.EUCentral1)
	}

	checkoutRegion := os.Getenv("CHECKOUT_REGION")
	if checkoutRegion == "true" {
		if region == string(connectivity.Hangzhou) {
			region = string(connectivity.EUCentral1)
			os.Setenv("ALICLOUD_REGION", region)
		}
	}
	backupRegionFind := false
	hangzhouRegionFind := false
	for _, r := range regions {
		if region == string(r) {
			find = true
		}
		if string(r) == backupRegion {
			backupRegionFind = true
		}
		if string(connectivity.Hangzhou) == string(r) {
			hangzhouRegionFind = true
		}
	}

	if (find && !supported) || (!find && supported) {
		if supported {
			if backupRegionFind {
				t.Logf("Skipping unsupported region %s. Supported regions: %s. Using %s as this test region", region, regions, backupRegion)
				os.Setenv("ALICLOUD_REGION", backupRegion)
				defaultRegionToTest = backupRegion
				return
			}
			if hangzhouRegionFind {
				t.Logf("Skipping unsupported region %s. Supported regions: %s. Using %s as this test region", region, regions, connectivity.Hangzhou)
				os.Setenv("ALICLOUD_REGION", string(connectivity.Hangzhou))
				defaultRegionToTest = string(connectivity.Hangzhou)
				return
			}
			if len(regions) > 0 {
				targetRegion := string(regions[0])
				t.Logf("Skipping unsupported region %s. Supported regions: %s. Using %s as this test region", region, regions, targetRegion)
				os.Setenv("ALICLOUD_REGION", targetRegion)
				defaultRegionToTest = targetRegion
				return
			}
			t.Skipf("Skipping unsupported region %s. Supported regions: %s.", region, regions)
		} else {
			if !backupRegionFind {
				t.Logf("Skipping unsupported region %s. Unsupported regions: %s. Using %s as this test region", region, regions, backupRegion)
				os.Setenv("ALICLOUD_REGION", backupRegion)
				defaultRegionToTest = backupRegion
				return
			}
			if !hangzhouRegionFind {
				t.Logf("Skipping unsupported region %s. Unsupported regions: %s. Using %s as this test region", region, regions, connectivity.Hangzhou)
				os.Setenv("ALICLOUD_REGION", string(connectivity.Hangzhou))
				defaultRegionToTest = string(connectivity.Hangzhou)
				return
			}
			t.Skipf("Skipping unsupported region %s. Unsupported regions: %s.", region, regions)
		}
		t.Skipped()
	}
}

// The instance member resource used by the ordering fixture is still owned by
// the root package, so its describe probe is mirrored here for the destroy check.
type cloudFirewallInstanceMemberTestService struct {
	client *connectivity.AliyunClient
}

func (s *cloudFirewallInstanceMemberTestService) DescribeCloudFirewallInstanceMember(id string) (object map[string]interface{}, err error) {
	var endpoint string
	var response map[string]interface{}
	action := "DescribeInstanceMembers"
	request := map[string]interface{}{
		"MemberUid": id,
	}

	wait := helper.IncrementalWait(3*time.Second, 3*time.Second)
	err = retry.Retry(5*time.Minute, func() *retry.RetryError {
		response, err = s.client.RpcPostWithEndpoint("Cloudfw", "2017-12-07", action, nil, request, true, endpoint)
		if err != nil {
			if helper.NeedRetry(err) {
				wait()
				return retry.RetryableError(err)
			} else if errs.IsExpectedErrors(err, []string{"not buy user"}) {
				endpoint = connectivity.CloudFirewallOpenAPIEndpointControlPolicy
				return retry.RetryableError(err)
			}
			return retry.NonRetryableError(err)
		}
		logs.AddDebug(action, response, request)
		return nil
	})
	if err != nil {
		if errs.IsExpectedErrors(err, []string{"InvalidResource.NotFound"}) {
			return object, errs.WrapErrorf(err, errs.NotFoundMsg, errs.AlibabaCloudSdkGoERROR)
		}
		return object, errs.WrapErrorf(err, errs.DefaultErrorMsg, id, action, errs.AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.Members", response)
	if err != nil {
		return object, errs.WrapErrorf(err, errs.FailedGetAttributeMsg, id, "$.Members", response)
	}
	members := v.([]interface{})
	if len(members) < 1 {
		return object, errs.WrapErrorf(err, errs.NotFoundMsg, errs.AlibabaCloudSdkGoERROR)
	}
	return members[0].(map[string]interface{}), nil
}

// Repeating addresses that already exist in state must not plan, while any
// membership or configured-order change must still plan.
func TestUnitCloudFirewallAddressBookDuplicateConvergence(t *testing.T) {
	a, b, c := "192.0.2.1/32", "198.51.100.1/32", "203.0.113.1/32"
	for _, tc := range []struct {
		name      string
		addresses []interface{}
		changed   bool
	}{
		{"duplicate tail", []interface{}{a, b, a}, false},
		{"duplicate head", []interface{}{a, a, b}, false},
		{"duplicate middle", []interface{}{a, b, b}, false},
		{"all duplicated", []interface{}{a, a, b, b, a}, false},
		{"duplicate with new member", []interface{}{a, a, c}, true},
		{"duplicate with removed member", []interface{}{a, a}, true},
		{"configured reorder", []interface{}{b, a, b}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := cloudfirewall.ResourceAddressBook()
			d := schema.TestResourceDataRaw(t, r.Schema, map[string]interface{}{"group_name": "test", "group_type": "ip", "description": "test", "address_list": []interface{}{a, b}})
			d.SetId("address-book")
			config := map[string]interface{}{"group_name": "test", "group_type": "ip", "description": "test", "address_list": tc.addresses}
			diff, err := r.Diff(context.Background(), d.State(), sdkterraform.NewResourceConfigRaw(config), nil)
			if err != nil {
				t.Fatal(err)
			}
			if changed := diff != nil && !diff.Empty(); changed != tc.changed {
				t.Fatalf("changed=%v want=%v diff=%#v", changed, tc.changed, diff)
			}
		})
	}

	// Without a configured address list the state mirrors the API view, including
	// its multiplicity. The plan must not rewrite such a computed list.
	r := cloudfirewall.ResourceAddressBook()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]interface{}{"group_name": "test", "group_type": "ip", "description": "test", "address_list": []interface{}{a, b, b}})
	d.SetId("address-book")
	diff, err := r.Diff(context.Background(), d.State(), sdkterraform.NewResourceConfigRaw(map[string]interface{}{"group_name": "test", "group_type": "ip", "description": "test"}), nil)
	if err != nil {
		t.Fatal(err)
	}
	if diff != nil && !diff.Empty() {
		t.Fatalf("an unconfigured address list must keep the API multiplicity, diff=%#v", diff)
	}
}

// Both write paths must submit one occurrence per address even when the
// configuration repeats them.
func TestUnitCloudFirewallAddressBookWritePathDedupes(t *testing.T) {
	a, b, c := "192.0.2.1/32", "198.51.100.1/32", "203.0.113.1/32"
	var created, updated []string
	client := cloudFirewallAddressBookTestClient(t, []interface{}{a, b}, func(req *http.Request) {
		if err := req.ParseForm(); err != nil {
			t.Error(err)
			return
		}
		switch req.Form.Get("Action") {
		case "AddAddressBook":
			created = append(created, req.Form.Get("AddressList"))
		case "ModifyAddressBook":
			updated = append(updated, req.Form.Get("AddressList"))
		}
	})
	r := cloudfirewall.ResourceAddressBook()

	config := map[string]interface{}{"group_name": "test", "group_type": "ip", "description": "test", "address_list": []interface{}{a, b, a}}
	diff, err := r.Diff(context.Background(), nil, sdkterraform.NewResourceConfigRaw(config), client)
	if err != nil {
		t.Fatal(err)
	}
	if diff == nil || diff.Empty() {
		t.Fatal("a new address book with configured addresses must plan")
	}
	state, diags := r.Apply(context.Background(), nil, diff, client)
	if diags.HasError() {
		t.Fatalf("%v", diags)
	}
	if len(created) != 1 || created[0] != a+","+b {
		t.Fatalf("AddAddressBook AddressList=%q want [%q]", created, a+","+b)
	}

	config["address_list"] = []interface{}{a, b, a, c}
	diff, err = r.Diff(context.Background(), state, sdkterraform.NewResourceConfigRaw(config), client)
	if err != nil {
		t.Fatal(err)
	}
	if diff == nil || diff.Empty() {
		t.Fatal("adding an address must plan")
	}
	if _, diags = r.Apply(context.Background(), state, diff, client); diags.HasError() {
		t.Fatalf("%v", diags)
	}
	if len(updated) != 1 || updated[0] != a+","+b+","+c {
		t.Fatalf("ModifyAddressBook AddressList=%q want [%q]", updated, a+","+b+","+c)
	}
}

// The duplicate fix must converge against the real API: the unchanged duplicate
// configuration plans empty after the first apply, and a reordered duplicate
// configuration still applies.
func TestAccAliCloudCloudFirewallAddressBook_addressListDuplicate(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	checkoutSupportedRegions(t, true, connectivity.CloudFirewallSupportRegions)
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return cloudfirewall.NewService(acctest.Provider.Meta().(*connectivity.AliyunClient))
	}, "DescribeAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testAcc%scloudfirewalladdressbook%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		CheckDestroy:             rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":   name,
					"group_type":   "ip",
					"description":  name,
					"address_list": []string{"10.21.0.0/16", "10.22.0.0/16", "10.21.0.0/16"},
				}),
				Check: testAccCheck(map[string]string{
					"address_list.#": "2",
				}),
			},
			{
				// The second apply of the same configuration must converge: the test
				// framework fails this step when the post-apply plan is not empty.
				Config: testAccConfig(map[string]interface{}{
					"address_list": []string{"10.21.0.0/16", "10.22.0.0/16", "10.21.0.0/16"},
				}),
				Check: testAccCheck(map[string]string{
					"address_list.#": "2",
				}),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"address_list": []string{"10.22.0.0/16", "10.21.0.0/16", "10.22.0.0/16"},
				}),
				Check: testAccCheck(map[string]string{
					"address_list.#": "2",
					"address_list.0": "10.22.0.0/16",
				}),
			},
		},
	})
}
