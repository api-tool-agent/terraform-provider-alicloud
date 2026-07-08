// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

func TestAccAlicloudApigGatewayDataSource(t *testing.T) {
	testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
	rand := acctest.RandIntRange(1000000, 9999999)

	idsConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudApigGatewaySourceConfig(rand, map[string]string{
			"ids": `["${alicloud_apig_gateway.default.id}"]`,
		}),
		fakeConfig: testAccCheckAlicloudApigGatewaySourceConfig(rand, map[string]string{
			"ids": `["${alicloud_apig_gateway.default.id}_fake"]`,
		}),
	}

	ResourceGroupIdConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudApigGatewaySourceConfig(rand, map[string]string{
			"ids":               `["${alicloud_apig_gateway.default.id}"]`,
			"resource_group_id": `"${data.alicloud_resource_manager_resource_groups.default.ids.0}"`,
		}),
		fakeConfig: testAccCheckAlicloudApigGatewaySourceConfig(rand, map[string]string{
			"ids":               `["${alicloud_apig_gateway.default.id}_fake"]`,
			"resource_group_id": `"${data.alicloud_resource_manager_resource_groups.default.ids.0}_fake"`,
		}),
	}
	GatewayNameConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudApigGatewaySourceConfig(rand, map[string]string{
			"ids":          `["${alicloud_apig_gateway.default.id}"]`,
			"gateway_name": `"${var.name}"`,
		}),
		fakeConfig: testAccCheckAlicloudApigGatewaySourceConfig(rand, map[string]string{
			"ids":          `["${alicloud_apig_gateway.default.id}_fake"]`,
			"gateway_name": `"${var.name}_fake"`,
		}),
	}

	allConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudApigGatewaySourceConfig(rand, map[string]string{
			"ids":               `["${alicloud_apig_gateway.default.id}"]`,
			"resource_group_id": `"${data.alicloud_resource_manager_resource_groups.default.ids.0}"`,

			"gateway_name": `"${var.name}"`,
		}),
		fakeConfig: testAccCheckAlicloudApigGatewaySourceConfig(rand, map[string]string{
			"ids":               `["${alicloud_apig_gateway.default.id}_fake"]`,
			"resource_group_id": `"${data.alicloud_resource_manager_resource_groups.default.ids.0}_fake"`,

			"gateway_name": `"${var.name}_fake"`,
		}),
	}

	ApigGatewayCheckInfo.dataSourceTestCheck(t, rand, idsConf, ResourceGroupIdConf, GatewayNameConf, allConf)
}

var existApigGatewayMapFunc = func(rand int) map[string]string {
	return map[string]string{
		"gateways.#":                    "1",
		"gateways.0.target_version":     CHECKSET,
		"gateways.0.record_total":       CHECKSET,
		"gateways.0.resource_group_id":  CHECKSET,
		"gateways.0.vpc.#":              CHECKSET,
		"gateways.0.gateway_type":       CHECKSET,
		"gateways.0.create_from":        CHECKSET,
		"gateways.0.zones.#":            CHECKSET,
		"gateways.0.version":            CHECKSET,
		"gateways.0.sub_domain_infos.#": CHECKSET,
		"gateways.0.payment_type":       CHECKSET,
		"gateways.0.tags.%":             CHECKSET,
		"gateways.0.status":             CHECKSET,
		"gateways.0.create_time":        CHECKSET,
		"gateways.0.gateway_edition":    CHECKSET,
		"gateways.0.load_balancers.#":   CHECKSET,
		"gateways.0.gateway_id":         CHECKSET,
		"gateways.0.security_group.#":   CHECKSET,
		"gateways.0.update_time":        CHECKSET,
		"gateways.0.gateway_name":       CHECKSET,
		"gateways.0.spec":               CHECKSET,
		"gateways.0.expire_time":        CHECKSET,
	}
}

var fakeApigGatewayMapFunc = func(rand int) map[string]string {
	return map[string]string{
		"gateways.#": "0",
	}
}

var ApigGatewayCheckInfo = dataSourceAttr{
	resourceId:   "data.alicloud_apig_gateways.default",
	existMapFunc: existApigGatewayMapFunc,
	fakeMapFunc:  fakeApigGatewayMapFunc,
}

func testAccCheckAlicloudApigGatewaySourceConfig(rand int, attrMap map[string]string) string {
	var pairs []string
	for k, v := range attrMap {
		pairs = append(pairs, k+" = "+v)
	}
	config := fmt.Sprintf(`
variable "name" {
	default = "tf-testAccApigGateway%d"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "gateway_vpc_pre3" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "gateway-test-vpc3"
}

resource "alicloud_vswitch" "gateway_vswitch_pre3" {
  is_default   = false
  vpc_id       = alicloud_vpc.gateway_vpc_pre3.id
  zone_id      = "cn-hangzhou-i"
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "gateway-test-vswitch3"
}



resource "alicloud_apig_gateway" "default" {
  network_access_config {
    type = "InternetAndIntranet"
  }
  resource_group_id = data.alicloud_resource_manager_resource_groups.default.ids.0
  vpc {
    vpc_id = alicloud_vpc.gateway_vpc_pre3.id
  }
  gateway_edition = "Professional"
  gateway_type    = "API"
  vswitch {
    vswitch_id = alicloud_vswitch.gateway_vswitch_pre3.id
  }
  zone_config {
    select_option = "Auto"
  }
  payment_type = "PayAsYouGo"
  gateway_name = "gateway-alt2-test"
  spec         = "apigw.small.x1"
  tags = {
    init-key = "init-value"
  }
}

data "alicloud_apig_gateways" "default" {
%s
}
`, rand, strings.Join(pairs, "\n   "))
	return config
}
