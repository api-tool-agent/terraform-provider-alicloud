// Package cms. This file is generated automatically. Please do not modify it manually, thank you!
package cms_test

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/acctest"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/service/cms"
	tfacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var checkCmsAlertRuleV2Destroy = func() resource.TestCheckFunc {
	var object map[string]interface{}
	rc := acctest.ResourceCheckInitWithDescribeMethod("alicloud_cms_alert_rule_v2.default", &object, func() interface{} {
		return cms.NewCmsServiceV2(acctest.Provider.Meta().(*connectivity.AliyunClient))
	}, "DescribeCmsAlertRuleV2")
	return rc.CheckResourceDestroy()
}()

// Test Cms AlertRuleV2. >>> Resource test cases, automatically generated.
// Case resource_AlertRuleV2_UMODEL_NoData_test 13068
func TestAccAliCloudCmsAlertRuleV2_basic13068(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic13068_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "content_template"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "UMODEL"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.0.enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_type", "instance"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "UMODEL_METRICSET_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_domain", "ecs"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric", "CPUUtilization"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric_set", "acs_ecs_dashboard"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-umodel-nodata"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "UMODEL_METRICSET_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold", "90"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic13068_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "content_template"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "UMODEL"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.0.enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_type", "instance"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "UMODEL_METRICSET_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_domain", "ecs"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric", "CPUUtilization"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric_set", "acs_ecs_dashboard"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-umodel-nodata-updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "UMODEL_METRICSET_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold", "90"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic13068_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "umodel no-data alert on $${metric}"

        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "UMODEL"
    }

        action_integration_config {
        enabled = false
    }

        arms_integration_config {
        enabled = false
    }

        query_config {
        entity_type = "instance"
        type = "UMODEL_METRICSET_QUERY"
        entity_domain = "ecs"
        metric = "CPUUtilization"
        metric_set = "acs_ecs_dashboard"
    }

        display_name = "regression-umodel-nodata"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "UMODEL_METRICSET_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
        threshold = 90
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic13068_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "umodel no-data alert on $${metric}"

        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "UMODEL"
    }

        action_integration_config {
        enabled = false
    }

        arms_integration_config {
        enabled = false
    }

        query_config {
        entity_type = "instance"
        type = "UMODEL_METRICSET_QUERY"
        entity_domain = "ecs"
        metric = "CPUUtilization"
        metric_set = "acs_ecs_dashboard"
    }

        display_name = "regression-umodel-nodata-updated"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "UMODEL_METRICSET_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
        threshold = 90
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_APM_Composite_test 12919
func TestAccAliCloudCmsAlertRuleV2_basic12919(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12919_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.errorrate"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-svc-composite-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-svc-composite-2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-composite-rewrite"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "SUM"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "5"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12919_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.errorrate"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-svc-composite-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-svc-composite-2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-composite-rewrite-updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "SUM"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "10"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12919_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.errorrate"
            group_by = ["rpc"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        service_id_list = ["arms-svc-composite-1", "arms-svc-composite-2"]
    }

        display_name = "regression-apm-composite-rewrite"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "SUM"
        threshold_list {
            severity = "WARNING"
            threshold = 5
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12919_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.errorrate"
            group_by = ["rpc"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        service_id_list = ["arms-svc-composite-1", "arms-svc-composite-2"]
    }

        display_name = "regression-apm-composite-rewrite-updated"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "SUM"
        threshold_list {
            severity = "CRITICAL"
            threshold = 10
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_Prometheus_NotifyStrategies_test 13054
func TestAccAliCloudCmsAlertRuleV2_basic13054(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic13054_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-strategies"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "NOTIFY_POLICY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.0", "strategy-1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic13054_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-strategies-m1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "NOTIFY_POLICY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.0", "strategy-id-003"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic13054_step2(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-strategies-m2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "NOTIFY_POLICY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.0", "strategy-4"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic13054_step3(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-strategies-m3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "NOTIFY_POLICY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.0", "strategy-5"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic13054_step4(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-strategies-m4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "NOTIFY_POLICY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.#", "0"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic13054_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 1"
    }

        display_name = "regression-notify-strategies"

        enabled = true

        notify_config {
        type = "NOTIFY_POLICY"
        notify_strategies = ["strategy-1"]
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic13054_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
    }

        display_name = "regression-notify-strategies-m1"

        enabled = true

        notify_config {
        type = "NOTIFY_POLICY"
        notify_strategies = ["strategy-id-003"]
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic13054_step2(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
    }

        display_name = "regression-notify-strategies-m2"

        enabled = true

        notify_config {
        type = "NOTIFY_POLICY"
        notify_strategies = ["strategy-4"]
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic13054_step3(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
    }

        display_name = "regression-notify-strategies-m3"

        enabled = true

        notify_config {
        type = "NOTIFY_POLICY"
        notify_strategies = ["strategy-5"]
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic13054_step4(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
    }

        display_name = "regression-notify-strategies-m4"

        enabled = true

        notify_config {
        type = "NOTIFY_POLICY"
        notify_strategies = []
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_ExpressEscalation_test 13053
func TestAccAliCloudCmsAlertRuleV2_basic13053(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic13053_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes < 0.2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-prom-express-escalation"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic13053_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes < 0.1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-prom-express-escalation-updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic13053_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes < 0.2"
    }

        display_name = "regression-prom-express-escalation"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic13053_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes < 0.1"
    }

        display_name = "regression-prom-express-escalation-updated"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_Prometheus_Basic_test 13052
func TestAccAliCloudCmsAlertRuleV2_basic13052(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic13052_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-basic-prometheus"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic13052_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-basic-prometheus-updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test-updated"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic13052_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 1"
    }

        display_name = "regression-basic-prometheus"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic13052_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
    }

        display_name = "regression-basic-prometheus-updated"

        enabled = false

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test-updated"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_Labels_Verify_test 13051
func TestAccAliCloudCmsAlertRuleV2_basic13051(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic13051_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "annotations.%", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "annotations.summary", "CPU usage high"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "annotations.description", "Node CPU is above 90%"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "cloudspec_labels_verify_test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "labels.%", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "labels.severity", "critical"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "labels.env", "test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic13051_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "annotations.%", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "annotations.summary", "CPU usage updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "annotations.description", "Updated CPU alert"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "annotations.runbook", "https://example.com"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "cloudspec_labels_verify_test_updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "labels.%", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "labels.severity", "warning"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "labels.team", "dev"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "labels.env", "staging"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic13051_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 1"
    }

        annotations = {
        "summary" = "CPU usage high"
        "description" = "Node CPU is above 90%%"
    }

        display_name = "cloudspec_labels_verify_test"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        labels = {
        "severity" = "critical"
        "env" = "test"
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic13051_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
    }

        annotations = {
        "summary" = "CPU usage updated"
        "description" = "Updated CPU alert"
        "runbook" = "https://example.com"
    }

        display_name = "cloudspec_labels_verify_test_updated"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        labels = {
        "severity" = "warning"
        "team" = "dev"
        "env" = "staging"
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_IntegrationConfig_test 12939
func TestAccAliCloudCmsAlertRuleV2_basic12939(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12939_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.0", "action-id-001"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.1", "action-id-002"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.2", "action-id-003"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.0.enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "integration-config-roundtrip-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12939_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.0", "action-id-004"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.0.enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "integration-config-roundtrip-test-s1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12939_step2(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.0", "action-id-005"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.1", "action-id-006"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.2", "action-id-007"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.0.enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "integration-config-roundtrip-test-s2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12939_step3(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.0", "action-id-008"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.1", "action-id-009"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.2", "action-id-010"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.0.enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "integration-config-roundtrip-test-s3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12939_step4(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.#", "0"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.0.enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "integration-config-roundtrip-test-s4-empty"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12939_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        action_integration_config {
        actions = ["action-id-001", "action-id-002", "action-id-003"]
        enabled = true
    }

        arms_integration_config {
        enabled = true
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 1"
    }

        display_name = "integration-config-roundtrip-test"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12939_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        action_integration_config {
        actions = ["action-id-004"]
        enabled = false
    }

        arms_integration_config {
        enabled = false
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
    }

        display_name = "integration-config-roundtrip-test-s1"

        enabled = false

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12939_step2(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        action_integration_config {
        actions = ["action-id-005", "action-id-006", "action-id-007"]
        enabled = true
    }

        arms_integration_config {
        enabled = true
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
    }

        display_name = "integration-config-roundtrip-test-s2"

        enabled = false

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12939_step3(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        action_integration_config {
        actions = ["action-id-008", "action-id-009", "action-id-010"]
        enabled = true
    }

        arms_integration_config {
        enabled = true
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 1"
    }

        display_name = "integration-config-roundtrip-test-s3"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12939_step4(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        action_integration_config {
        actions = []
        enabled = true
    }

        arms_integration_config {
        enabled = true
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 1"
    }

        display_name = "integration-config-roundtrip-test-s4-empty"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_APM_Aggregate_test 12937
func TestAccAliCloudCmsAlertRuleV2_basic12937(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12937_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.rt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.1", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.value", "http"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-agg-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-agg-2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-aggregate"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "SUM"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "1000"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.1.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.1.threshold", "500"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12937_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.rt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.1", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.value", "http"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-agg-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-agg-2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-aggregate-max"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "LT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "MAX"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "200"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12937_step2(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.rt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.value", "http"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-agg-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-agg-2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-aggregate-p50"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "P50"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "INFO"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "50"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12937_step3(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.rt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.value", "http"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-agg-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-agg-2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-aggregate-p75"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "P75"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "ERROR"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "75"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12937_step4(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.rt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.1", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.value", "http"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-agg-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-agg-2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-aggregate-p90"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "LT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "P90"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "90"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12937_step5(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.rt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.1", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.value", "http"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-agg-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-agg-2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-aggregate-p99"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "P99"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "99"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12937_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.rt"
            group_by = ["rpc", "rpcType"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        filter_list {
            type = "EQ"
            value = "http"
            key = "rpcType"
        }
        service_id_list = ["arms-agg-1", "arms-agg-2"]
    }

        display_name = "regression-apm-aggregate"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "SUM"
        threshold_list {
            severity = "CRITICAL"
            threshold = 1000
        }
        threshold_list {
            severity = "WARNING"
            threshold = 500
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12937_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.rt"
            group_by = ["rpc", "rpcType"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        filter_list {
            type = "EQ"
            value = "http"
            key = "rpcType"
        }
        service_id_list = ["arms-agg-1", "arms-agg-2"]
    }

        display_name = "regression-apm-aggregate-max"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "LT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "MAX"
        threshold_list {
            severity = "WARNING"
            threshold = 200
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12937_step2(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.rt"
            group_by = ["rpc"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        filter_list {
            type = "EQ"
            value = "http"
            key = "rpcType"
        }
        service_id_list = ["arms-agg-1", "arms-agg-2"]
    }

        display_name = "regression-apm-aggregate-p50"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "P50"
        threshold_list {
            severity = "INFO"
            threshold = 50
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12937_step3(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.rt"
            group_by = ["rpcType"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        filter_list {
            type = "EQ"
            value = "http"
            key = "rpcType"
        }
        service_id_list = ["arms-agg-1", "arms-agg-2"]
    }

        display_name = "regression-apm-aggregate-p75"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "P75"
        threshold_list {
            severity = "ERROR"
            threshold = 75
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12937_step4(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.rt"
            group_by = ["rpc", "rpcType"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        filter_list {
            type = "EQ"
            value = "http"
            key = "rpcType"
        }
        service_id_list = ["arms-agg-1", "arms-agg-2"]
    }

        display_name = "regression-apm-aggregate-p90"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "LT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "P90"
        threshold_list {
            severity = "WARNING"
            threshold = 90
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12937_step5(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.rt"
            group_by = ["rpc", "rpcType"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        filter_list {
            type = "EQ"
            value = "http"
            key = "rpcType"
        }
        service_id_list = ["arms-agg-1", "arms-agg-2"]
    }

        display_name = "regression-apm-aggregate-p99"

        enabled = false

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "P99"
        threshold_list {
            severity = "CRITICAL"
            threshold = 99
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_QueryConfig_Advanced_test 12936
func TestAccAliCloudCmsAlertRuleV2_basic12936(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12936_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "UMODEL"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_type", "instance"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "UMODEL_METRICSET_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_domain", "ecs"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.0.operator", "="),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.0.field", "regionId"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.0.value", "cn-hangzhou"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.1.operator", "="),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.1.field", "instanceId"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.1.value", "i-bp1abc123def"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric", "CPUUtilization"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric_set", "acs_ecs_dashboard"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_fields.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_fields.0.field", "name"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_fields.0.value", "my-instance"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-queryconfig-advanced"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "UMODEL_METRICSET_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold", "90"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12936_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "UMODEL"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_type", "instance"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "UMODEL_METRICSET_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_domain", "ecs"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.0.operator", "="),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.0.field", "regionId"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.0.value", "cn-shanghai"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.1.operator", "="),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.1.field", "instanceId"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.1.value", "i-bp2ghi456jkl"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.2.operator", "="),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.2.field", "status"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_filters.2.value", "Running"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric", "memory_usedutilization"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric_set", "acs_ecs_dashboard"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_fields.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_fields.0.field", "name"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_fields.0.value", "my-ecs-instance"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_fields.1.field", "hostName"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_fields.1.value", "web-server-01"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-queryconfig-advanced-updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test-updated"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "UMODEL_METRICSET_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold", "95"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12936_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "UMODEL"
    }

        query_config {
        entity_type = "instance"
        type = "UMODEL_METRICSET_QUERY"
        entity_domain = "ecs"
        entity_filters {
            operator = "="
            field = "regionId"
            value = "cn-hangzhou"
        }
        entity_filters {
            operator = "="
            field = "instanceId"
            value = "i-bp1abc123def"
        }
        metric = "CPUUtilization"
        metric_set = "acs_ecs_dashboard"
        entity_fields {
            field = "name"
            value = "my-instance"
        }
    }

        display_name = "regression-queryconfig-advanced"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "UMODEL_METRICSET_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
        threshold = 90
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12936_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "UMODEL"
    }

        query_config {
        entity_type = "instance"
        type = "UMODEL_METRICSET_QUERY"
        entity_domain = "ecs"
        entity_filters {
            operator = "="
            field = "regionId"
            value = "cn-shanghai"
        }
        entity_filters {
            operator = "="
            field = "instanceId"
            value = "i-bp2ghi456jkl"
        }
        entity_filters {
            operator = "="
            field = "status"
            value = "Running"
        }
        metric = "memory_usedutilization"
        metric_set = "acs_ecs_dashboard"
        entity_fields {
            field = "name"
            value = "my-ecs-instance"
        }
        entity_fields {
            field = "hostName"
            value = "web-server-01"
        }
    }

        display_name = "regression-queryconfig-advanced-updated"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test-updated"]
        }
    }

        condition_config {
        operator = "GT"
        type = "UMODEL_METRICSET_CONDITION"
        severity = "WARNING"
        duration_secs = 120
        threshold = 95
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_APM_CompareList_test 12935
func TestAccAliCloudCmsAlertRuleV2_basic12935(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12935_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.rt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-svc-compare-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-svc-compare-2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-comparelist"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.compare_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.compare_list.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.compare_list.0.aggregate", "AVG"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.compare_list.0.yoy_time_value", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.compare_list.0.threshold", "80"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.compare_list.0.yoy_time_unit", "hour"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_COMPOSITE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.relation", "OR"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12935_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.rt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-svc-compare-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-svc-compare-2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-comparelist-updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.compare_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.compare_list.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.compare_list.0.aggregate", "P99"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.compare_list.0.yoy_time_value", "24"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.compare_list.0.threshold", "95"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.compare_list.0.yoy_time_unit", "hour"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_COMPOSITE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.relation", "AND"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12935_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.rt"
            group_by = ["rpc"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        service_id_list = ["arms-svc-compare-1", "arms-svc-compare-2"]
    }

        display_name = "regression-apm-comparelist"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        compare_list {
            operator = "GT"
            aggregate = "AVG"
            yoy_time_value = 1
            threshold = 80
            yoy_time_unit = "hour"
        }
        type = "APM_COMPOSITE_CONDITION"
        relation = "OR"
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12935_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.rt"
            group_by = ["rpc"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        service_id_list = ["arms-svc-compare-1", "arms-svc-compare-2"]
    }

        display_name = "regression-apm-comparelist-updated"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        compare_list {
            operator = "GT"
            aggregate = "P99"
            yoy_time_value = 24
            threshold = 95
            yoy_time_unit = "hour"
        }
        type = "APM_COMPOSITE_CONDITION"
        relation = "AND"
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_APM_Operator_test 12934
func TestAccAliCloudCmsAlertRuleV2_basic12934(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12934_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.count"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.value", "http"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-app-001"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-app-002"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-operator"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "COUNT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "1000"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.1.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.1.threshold", "500"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12934_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.count"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.value", "http"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-app-001"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-app-002"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-operator-gte"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "MIN"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "800"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12934_step2(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.count"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.value", "http"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-app-001"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-app-002"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-operator-lt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "LT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "CONTINUES"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "INFO"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "10"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12934_step3(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.count"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.value", "http"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-app-001"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-app-002"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-operator-lte"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "LT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "COUNT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "20"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12934_step4(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.count"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.value", "http"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-app-001"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-app-002"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-operator-eq"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "MIN"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "ERROR"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "0"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12934_step5(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.count"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.value", "http"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-app-001"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-app-002"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-operator-ne"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "NE"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "CONTINUES"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "999"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12934_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.count"
            group_by = ["rpc"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        filter_list {
            type = "EQ"
            value = "http"
            key = "rpcType"
        }
        service_id_list = ["arms-app-001", "arms-app-002"]
    }

        display_name = "regression-apm-operator"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "COUNT"
        threshold_list {
            severity = "CRITICAL"
            threshold = 1000
        }
        threshold_list {
            severity = "WARNING"
            threshold = 500
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12934_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.count"
            group_by = ["rpc"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        filter_list {
            type = "EQ"
            value = "http"
            key = "rpcType"
        }
        service_id_list = ["arms-app-001", "arms-app-002"]
    }

        display_name = "regression-apm-operator-gte"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "MIN"
        threshold_list {
            severity = "WARNING"
            threshold = 800
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12934_step2(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.count"
            group_by = ["rpc"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        filter_list {
            type = "EQ"
            value = "http"
            key = "rpcType"
        }
        service_id_list = ["arms-app-001", "arms-app-002"]
    }

        display_name = "regression-apm-operator-lt"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "LT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "CONTINUES"
        threshold_list {
            severity = "INFO"
            threshold = 10
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12934_step3(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.count"
            group_by = ["rpc"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        filter_list {
            type = "EQ"
            value = "http"
            key = "rpcType"
        }
        service_id_list = ["arms-app-001", "arms-app-002"]
    }

        display_name = "regression-apm-operator-lte"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "LT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "COUNT"
        threshold_list {
            severity = "WARNING"
            threshold = 20
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12934_step4(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.count"
            group_by = ["rpc"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        filter_list {
            type = "EQ"
            value = "http"
            key = "rpcType"
        }
        service_id_list = ["arms-app-001", "arms-app-002"]
    }

        display_name = "regression-apm-operator-eq"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "EQ"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "MIN"
        threshold_list {
            severity = "ERROR"
            threshold = 0
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12934_step5(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.count"
            group_by = ["rpc"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        filter_list {
            type = "EQ"
            value = "http"
            key = "rpcType"
        }
        service_id_list = ["arms-app-001", "arms-app-002"]
    }

        display_name = "regression-apm-operator-ne"

        enabled = false

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "NE"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "CONTINUES"
        threshold_list {
            severity = "CRITICAL"
            threshold = 999
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_SimpleEscalation_test 12932
func TestAccAliCloudCmsAlertRuleV2_basic12932(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12932_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "UMODEL"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_type", "instance"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "UMODEL_METRICSET_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_domain", "ecs"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric", "memory_usedutilization"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric_set", "acs_ecs_dashboard"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-umodel-simple-escalation"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GE"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "UMODEL_METRICSET_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold", "20"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12932_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "UMODEL"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_type", "instance"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "UMODEL_METRICSET_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_domain", "ecs"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric", "memory_usedutilization"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric_set", "acs_ecs_dashboard"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-umodel-simple-escalation-updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "LE"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "UMODEL_METRICSET_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "180"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold", "90"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12932_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "UMODEL"
    }

        query_config {
        entity_type = "instance"
        type = "UMODEL_METRICSET_QUERY"
        entity_domain = "ecs"
        metric = "memory_usedutilization"
        metric_set = "acs_ecs_dashboard"
    }

        display_name = "regression-umodel-simple-escalation"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GE"
        type = "UMODEL_METRICSET_CONDITION"
        severity = "WARNING"
        duration_secs = 120
        threshold = 20
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12932_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "UMODEL"
    }

        query_config {
        entity_type = "instance"
        type = "UMODEL_METRICSET_QUERY"
        entity_domain = "ecs"
        metric = "memory_usedutilization"
        metric_set = "acs_ecs_dashboard"
    }

        display_name = "regression-umodel-simple-escalation-updated"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "LE"
        type = "UMODEL_METRICSET_CONDITION"
        severity = "CRITICAL"
        duration_secs = 180
        threshold = 90
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_APM_Simple_test 12930
func TestAccAliCloudCmsAlertRuleV2_basic12930(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12930_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.rt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.1", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.value", "http"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "NE"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.value", "/health"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "svc-001"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "svc-002"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-simple-8"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "AVG"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "1000"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.1.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.1.threshold", "500"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12930_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.rt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.1", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "120"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.value", "/orders"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "svc-001"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "svc-003"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.2", "svc-004"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-simple-8-modified"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "AVG"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "50"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.1.severity", "INFO"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.1.threshold", "30"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.2.severity", "ERROR"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.2.threshold", "80"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12930_step2(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.rt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "120"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.value", "/orders"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "svc-009"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-simple-8-modified-2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "AVG"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "50"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.1.severity", "INFO"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.1.threshold", "30"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.2.severity", "ERROR"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.2.threshold", "80"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12930_step3(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.rt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpcType"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.1", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "120"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.value", "/orders"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "svc-010"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "svc-011"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-simple-8-modified-3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "AVG"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "50"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.1.severity", "INFO"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.1.threshold", "30"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.2.severity", "ERROR"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.2.threshold", "80"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12930_step4(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.rt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "0"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "120"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "EQ"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.value", "/orders"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "svc-012"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-simple-8-modified-4-empty"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "AVG"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "50"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.1.severity", "INFO"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.1.threshold", "30"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.2.severity", "ERROR"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.2.threshold", "80"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12930_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.rt"
            group_by = ["rpc", "rpcType"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "EQ"
            value = "http"
            key = "rpcType"
        }
        filter_list {
            type = "NE"
            value = "/health"
            key = "rpc"
        }
        service_id_list = ["svc-001", "svc-002"]
    }

        display_name = "regression-apm-simple-8"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "AVG"
        threshold_list {
            severity = "CRITICAL"
            threshold = 1000
        }
        threshold_list {
            severity = "WARNING"
            threshold = 500
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12930_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.rt"
            group_by = ["rpc", "rpcType"]
            window_secs = 120
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "EQ"
            value = "/orders"
            key = "rpc"
        }
        service_id_list = ["svc-001", "svc-003", "svc-004"]
    }

        display_name = "regression-apm-simple-8-modified"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "AVG"
        threshold_list {
            severity = "WARNING"
            threshold = 50
        }
        threshold_list {
            severity = "INFO"
            threshold = 30
        }
        threshold_list {
            severity = "ERROR"
            threshold = 80
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12930_step2(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.rt"
            group_by = ["rpc"]
            window_secs = 120
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "EQ"
            value = "/orders"
            key = "rpc"
        }
        service_id_list = ["svc-009"]
    }

        display_name = "regression-apm-simple-8-modified-2"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "AVG"
        threshold_list {
            severity = "WARNING"
            threshold = 50
        }
        threshold_list {
            severity = "INFO"
            threshold = 30
        }
        threshold_list {
            severity = "ERROR"
            threshold = 80
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12930_step3(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.rt"
            group_by = ["rpcType", "rpc"]
            window_secs = 120
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "EQ"
            value = "/orders"
            key = "rpc"
        }
        service_id_list = ["svc-010", "svc-011"]
    }

        display_name = "regression-apm-simple-8-modified-3"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "AVG"
        threshold_list {
            severity = "WARNING"
            threshold = 50
        }
        threshold_list {
            severity = "INFO"
            threshold = 30
        }
        threshold_list {
            severity = "ERROR"
            threshold = 80
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12930_step4(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.transaction.rt"
            group_by = []
            window_secs = 120
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "EQ"
            value = "/orders"
            key = "rpc"
        }
        service_id_list = ["svc-012"]
    }

        display_name = "regression-apm-simple-8-modified-4-empty"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "AVG"
        threshold_list {
            severity = "WARNING"
            threshold = 50
        }
        threshold_list {
            severity = "INFO"
            threshold = 30
        }
        threshold_list {
            severity = "ERROR"
            threshold = 80
        }
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_UMODEL_test 12929
func TestAccAliCloudCmsAlertRuleV2_basic12929(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12929_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "content_template"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "UMODEL"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.0.enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_type", "instance"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "UMODEL_METRICSET_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_domain", "ecs"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric", "CPUUtilization"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.0.operator", "="),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.0.value", "web-server"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.0.name", "app"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.1.operator", "="),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.1.value", "production"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.1.name", "env"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric_set", "acs_ecs_dashboard"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-umodel-10"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "UMODEL_METRICSET_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold", "90"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12929_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "content_template"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "UMODEL"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.0", "action-integration-001"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.0.enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_type", "instance"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "UMODEL_METRICSET_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_domain", "ecs"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric", "CPUUtilization"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.0.operator", "="),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.0.value", "api-gateway"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.0.name", "app"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.1.operator", "="),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.1.value", "staging"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.1.name", "env"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.2.operator", "="),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.2.value", "cn-hangzhou"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.label_filters.2.name", "region"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric_set", "acs_ecs_dashboard"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-umodel-10-updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "UMODEL_METRICSET_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold", "95"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12929_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "umodel test alert on $${metric}"

        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "UMODEL"
    }

        action_integration_config {
        enabled = false
    }

        arms_integration_config {
        enabled = false
    }

        query_config {
        entity_type = "instance"
        type = "UMODEL_METRICSET_QUERY"
        entity_domain = "ecs"
        metric = "CPUUtilization"
        label_filters {
            operator = "="
            value = "web-server"
            name = "app"
        }
        label_filters {
            operator = "="
            value = "production"
            name = "env"
        }
        metric_set = "acs_ecs_dashboard"
    }

        display_name = "regression-umodel-10"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "UMODEL_METRICSET_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
        threshold = 90
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12929_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "umodel test alert updated on $${metric}"

        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "UMODEL"
    }

        action_integration_config {
        actions = ["action-integration-001"]
        enabled = true
    }

        arms_integration_config {
        enabled = true
    }

        query_config {
        entity_type = "instance"
        type = "UMODEL_METRICSET_QUERY"
        entity_domain = "ecs"
        metric = "CPUUtilization"
        label_filters {
            operator = "="
            value = "api-gateway"
            name = "app"
        }
        label_filters {
            operator = "="
            value = "staging"
            name = "env"
        }
        label_filters {
            operator = "="
            value = "cn-hangzhou"
            name = "region"
        }
        metric_set = "acs_ecs_dashboard"
    }

        display_name = "regression-umodel-10-updated"

        enabled = false

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "UMODEL_METRICSET_CONDITION"
        severity = "CRITICAL"
        duration_secs = 120
        threshold = 95
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_Prometheus_Rich_test 12928
func TestAccAliCloudCmsAlertRuleV2_basic12928(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12928_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "content_template"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "rate(http_requests_total[5m]) > 100"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-enriched-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.utc_offset", "+08:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.#", "7"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.0", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.1", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.2", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.3", "4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.4", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.5", "6"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.6", "7"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_end_time", "23:59"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.silence_time_secs", "86400"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_start_time", "00:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "group-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.1", "group-2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.2", "group-3"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12928_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "content_template"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "rate(http_requests_total[5m]) > 200"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-enriched-1-updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.utc_offset", "+08:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.#", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.0", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.1", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.2", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.3", "4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.4", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_end_time", "18:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.silence_time_secs", "43200"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_start_time", "09:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test-updated"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12928_step2(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "content_template"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "rate(http_requests_total[5m]) > 200"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-enriched-1-m2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.utc_offset", "+08:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.#", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.0", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.1", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.2", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.3", "4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.4", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_end_time", "18:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.silence_time_secs", "43200"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_start_time", "09:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "group-4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.1", "group-5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.2", "group-6"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12928_step3(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "content_template"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "rate(http_requests_total[5m]) > 200"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-enriched-1-m3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.utc_offset", "+08:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.#", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.0", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.1", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.2", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.3", "4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.4", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_end_time", "18:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.silence_time_secs", "43200"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_start_time", "09:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "group-4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.1", "group-5"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12928_step4(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "content_template"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "rate(http_requests_total[5m]) > 200"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-enriched-1-m4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.utc_offset", "+08:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.#", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.0", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.1", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.2", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.3", "4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.4", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_end_time", "18:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.silence_time_secs", "43200"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_start_time", "09:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "group-7"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12928_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "告警：$${alertName} 触发"

        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "rate(http_requests_total[5m]) > 100"
    }

        display_name = "regression-enriched-1"

        enabled = true

        notify_config {
        utc_offset = "+08:00"
        type = "DIRECT_NOTIFY"
        active_days = [1, 2, 3, 4, 5, 6, 7]
        active_end_time = "23:59"
        silence_time_secs = 86400
        active_start_time = "00:00"
        channels {
            type = "GROUP"
            identifiers = ["group-1", "group-2", "group-3"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12928_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "更新告警：$${alertName}"

        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "rate(http_requests_total[5m]) > 200"
    }

        display_name = "regression-enriched-1-updated"

        enabled = true

        notify_config {
        utc_offset = "+08:00"
        type = "DIRECT_NOTIFY"
        active_days = [1, 2, 3, 4, 5]
        active_end_time = "18:00"
        silence_time_secs = 43200
        active_start_time = "09:00"
        channels {
            type = "GROUP"
            identifiers = ["regression-test-updated"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12928_step2(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "更新告警：$${alertName}"

        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "rate(http_requests_total[5m]) > 200"
    }

        display_name = "regression-enriched-1-m2"

        enabled = true

        notify_config {
        utc_offset = "+08:00"
        type = "DIRECT_NOTIFY"
        active_days = [1, 2, 3, 4, 5]
        active_end_time = "18:00"
        silence_time_secs = 43200
        active_start_time = "09:00"
        channels {
            type = "GROUP"
            identifiers = ["group-4", "group-5", "group-6"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12928_step3(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "更新告警：$${alertName}"

        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "rate(http_requests_total[5m]) > 200"
    }

        display_name = "regression-enriched-1-m3"

        enabled = true

        notify_config {
        utc_offset = "+08:00"
        type = "DIRECT_NOTIFY"
        active_days = [1, 2, 3, 4, 5]
        active_end_time = "18:00"
        silence_time_secs = 43200
        active_start_time = "09:00"
        channels {
            type = "GROUP"
            identifiers = ["group-4", "group-5"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12928_step4(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "更新告警：$${alertName}"

        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "rate(http_requests_total[5m]) > 200"
    }

        display_name = "regression-enriched-1-m4"

        enabled = true

        notify_config {
        utc_offset = "+08:00"
        type = "DIRECT_NOTIFY"
        active_days = [1, 2, 3, 4, 5]
        active_end_time = "18:00"
        silence_time_secs = 43200
        active_start_time = "09:00"
        channels {
            type = "GROUP"
            identifiers = ["group-7"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_Prometheus_Array_test 12924
func TestAccAliCloudCmsAlertRuleV2_basic12924(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12924_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-severity-array"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.utc_offset", "+08:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.#", "7"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.0", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.1", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.2", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.3", "4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.4", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.5", "6"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.6", "7"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_end_time", "23:59"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_start_time", "00:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.silence_time_secs", "86400"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "group-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.1", "group-2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.2", "group-3"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "INFO"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12924_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-severity-array-updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.utc_offset", "+08:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.#", "6"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.0", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.1", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.2", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.3", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.4", "6"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.5", "7"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_end_time", "20:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_start_time", "08:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.silence_time_secs", "43200"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "group-a"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.1", "group-b"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "ERROR"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12924_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 1"
    }

        display_name = "regression-severity-array"

        enabled = true

        notify_config {
        utc_offset = "+08:00"
        type = "DIRECT_NOTIFY"
        active_days = [1, 2, 3, 4, 5, 6, 7]
        active_end_time = "23:59"
        active_start_time = "00:00"
        silence_time_secs = 86400
        channels {
            type = "GROUP"
            identifiers = ["group-1", "group-2", "group-3"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "INFO"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12924_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
    }

        display_name = "regression-severity-array-updated"

        enabled = true

        notify_config {
        utc_offset = "+08:00"
        type = "DIRECT_NOTIFY"
        active_days = [1, 2, 3, 5, 6, 7]
        active_end_time = "20:00"
        active_start_time = "08:00"
        silence_time_secs = 43200
        channels {
            type = "GROUP"
            identifiers = ["group-a", "group-b"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "ERROR"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_NotifyChannels_test 12923
func TestAccAliCloudCmsAlertRuleV2_basic12923(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12923_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "cmsxx-prometheus-1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up == 0"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-channels"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.utc_offset", "+08:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.0", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.1", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.2", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_end_time", "23:59"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.silence_time_secs", "86400"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_start_time", "00:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "cms-alarm-group-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.1", "cms-alarm-group-2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.type", "DINGTALK"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.0", "dingtalk-webhook-url-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.1", "dingtalk-webhook-url-2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.type", "WEBHOOK"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.identifiers.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.identifiers.0", "https://webhook.example.com/alert"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.identifiers.1", "https://webhook2.example.com/alert"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12923_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "cmsxx-prometheus-1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up == 0"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-channels-m1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.utc_offset", "+08:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.#", "7"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.0", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.1", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.2", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.3", "4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.4", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.5", "6"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.6", "7"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_end_time", "23:59"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.silence_time_secs", "86400"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_start_time", "00:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "CONTACT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "contact-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.1", "contact-2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.2", "contact-3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.type", "FEISHU"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.0", "feishu-webhook-url-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.1", "feishu-webhook-url-2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.type", "SLACK"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.identifiers.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.identifiers.0", "slack-channel-url-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.identifiers.1", "slack-channel-url-2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.3.type", "WEIXIN"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.3.identifiers.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.3.identifiers.0", "weixin-bot-key-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.3.identifiers.1", "weixin-bot-key-2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12923_step2(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "cmsxx-prometheus-1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up == 0"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-channels-m2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.utc_offset", "+08:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.0", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.1", "4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.2", "6"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_end_time", "20:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.silence_time_secs", "43200"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_start_time", "08:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "group-a"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.1", "group-b"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.2", "group-c"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.type", "DINGTALK"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.0", "dt-url-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.1", "dt-url-2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.2", "dt-url-3"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12923_step3(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "cmsxx-prometheus-1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up == 0"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-channels-m3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.utc_offset", "+08:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.#", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.0", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.1", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.2", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.3", "4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.4", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_end_time", "18:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.silence_time_secs", "86400"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_start_time", "09:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "CONTACT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "admin-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.1", "admin-2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.type", "WEBHOOK"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.0", "https://hook1.example.com"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.1", "https://hook2.example.com"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12923_step4(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "cmsxx-prometheus-1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up == 0"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-channels-m4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.utc_offset", "+08:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.#", "7"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.0", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.1", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.2", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.3", "4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.4", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.5", "6"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.6", "7"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_end_time", "23:59"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.silence_time_secs", "86400"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_start_time", "00:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "ops-group-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.1", "ops-group-2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.type", "DINGTALK"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.0", "dt-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.1", "dt-2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.type", "FEISHU"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.identifiers.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.identifiers.0", "feishu-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.identifiers.1", "feishu-2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12923_step5(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "90"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "cmsxx-prometheus-1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up == 0"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-channels-m5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.utc_offset", "+08:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.#", "4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.0", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.1", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.2", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_days.3", "7"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_end_time", "22:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.silence_time_secs", "86400"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.active_start_time", "06:00"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "CONTACT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "dev-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.1", "dev-2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.2", "dev-3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.type", "SLACK"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.0", "slack-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.1.identifiers.1", "slack-2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.type", "WEIXIN"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.identifiers.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.identifiers.0", "wx-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.identifiers.1", "wx-2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.2.identifiers.2", "wx-3"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "90"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12923_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "cmsxx-prometheus-1"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up == 0"
    }

        display_name = "regression-notify-channels"

        enabled = true

        notify_config {
        utc_offset = "+08:00"
        type = "DIRECT_NOTIFY"
        active_days = [1, 3, 5]
        active_end_time = "23:59"
        silence_time_secs = 86400
        active_start_time = "00:00"
        channels {
            type = "GROUP"
            identifiers = ["cms-alarm-group-1", "cms-alarm-group-2"]
        }
        channels {
            type = "DINGTALK"
            identifiers = ["dingtalk-webhook-url-1", "dingtalk-webhook-url-2"]
        }
        channels {
            type = "WEBHOOK"
            identifiers = ["https://webhook.example.com/alert", "https://webhook2.example.com/alert"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12923_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "cmsxx-prometheus-1"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up == 0"
    }

        display_name = "regression-notify-channels-m1"

        enabled = true

        notify_config {
        utc_offset = "+08:00"
        type = "DIRECT_NOTIFY"
        active_days = [1, 2, 3, 4, 5, 6, 7]
        active_end_time = "23:59"
        silence_time_secs = 86400
        active_start_time = "00:00"
        channels {
            type = "CONTACT"
            identifiers = ["contact-1", "contact-2", "contact-3"]
        }
        channels {
            type = "FEISHU"
            identifiers = ["feishu-webhook-url-1", "feishu-webhook-url-2"]
        }
        channels {
            type = "SLACK"
            identifiers = ["slack-channel-url-1", "slack-channel-url-2"]
        }
        channels {
            type = "WEIXIN"
            identifiers = ["weixin-bot-key-1", "weixin-bot-key-2"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12923_step2(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "cmsxx-prometheus-1"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up == 0"
    }

        display_name = "regression-notify-channels-m2"

        enabled = true

        notify_config {
        utc_offset = "+08:00"
        type = "DIRECT_NOTIFY"
        active_days = [2, 4, 6]
        active_end_time = "20:00"
        silence_time_secs = 43200
        active_start_time = "08:00"
        channels {
            type = "GROUP"
            identifiers = ["group-a", "group-b", "group-c"]
        }
        channels {
            type = "DINGTALK"
            identifiers = ["dt-url-1", "dt-url-2", "dt-url-3"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12923_step3(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "cmsxx-prometheus-1"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up == 0"
    }

        display_name = "regression-notify-channels-m3"

        enabled = true

        notify_config {
        utc_offset = "+08:00"
        type = "DIRECT_NOTIFY"
        active_days = [1, 2, 3, 4, 5]
        active_end_time = "18:00"
        silence_time_secs = 86400
        active_start_time = "09:00"
        channels {
            type = "CONTACT"
            identifiers = ["admin-1", "admin-2"]
        }
        channels {
            type = "WEBHOOK"
            identifiers = ["https://hook1.example.com", "https://hook2.example.com"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12923_step4(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "cmsxx-prometheus-1"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up == 0"
    }

        display_name = "regression-notify-channels-m4"

        enabled = true

        notify_config {
        utc_offset = "+08:00"
        type = "DIRECT_NOTIFY"
        active_days = [1, 2, 3, 4, 5, 6, 7]
        active_end_time = "23:59"
        silence_time_secs = 86400
        active_start_time = "00:00"
        channels {
            type = "GROUP"
            identifiers = ["ops-group-1", "ops-group-2"]
        }
        channels {
            type = "DINGTALK"
            identifiers = ["dt-1", "dt-2"]
        }
        channels {
            type = "FEISHU"
            identifiers = ["feishu-1", "feishu-2"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12923_step5(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 90
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "cmsxx-prometheus-1"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up == 0"
    }

        display_name = "regression-notify-channels-m5"

        enabled = true

        notify_config {
        utc_offset = "+08:00"
        type = "DIRECT_NOTIFY"
        active_days = [1, 3, 5, 7]
        active_end_time = "22:00"
        silence_time_secs = 86400
        active_start_time = "06:00"
        channels {
            type = "CONTACT"
            identifiers = ["dev-1", "dev-2", "dev-3"]
        }
        channels {
            type = "SLACK"
            identifiers = ["slack-1", "slack-2"]
        }
        channels {
            type = "WEIXIN"
            identifiers = ["wx-1", "wx-2", "wx-3"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 90
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_Prometheus_DataCheck_test 12922
func TestAccAliCloudCmsAlertRuleV2_basic12922(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12922_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.enable_data_complete_check", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-data-check"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12922_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.enable_data_complete_check", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-data-check-updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test-updated"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12922_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 1"
        enable_data_complete_check = true
    }

        display_name = "regression-data-check"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12922_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
        enable_data_complete_check = false
    }

        display_name = "regression-data-check-updated"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test-updated"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_Prometheus_Condition_test 12921
func TestAccAliCloudCmsAlertRuleV2_basic12921(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12921_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "rate(http_errors_total[5m]) / rate(http_requests_total[5m]) > 0.05"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-prom-condition-rewrite"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12921_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "rate(http_errors_total[10m]) / rate(http_requests_total[10m]) > 0.1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-prom-condition-rewrite-updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "180"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12921_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "rate(http_errors_total[5m]) / rate(http_requests_total[5m]) > 0.05"
    }

        display_name = "regression-prom-condition-rewrite"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12921_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "rate(http_errors_total[10m]) / rate(http_requests_total[10m]) > 0.1"
    }

        display_name = "regression-prom-condition-rewrite-updated"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 180
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_EmptyArray_Repro_test 12920
func TestAccAliCloudCmsAlertRuleV2_basic12920(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12920_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "content_template", "empty array repro test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "UMODEL"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.0", "action-001"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.1", "action-002"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.0.enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_type", "instance"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "UMODEL_METRICSET_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_domain", "ecs"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric", "CPUUtilization"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric_set", "acs_ecs_dashboard"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "empty-array-repro-init"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "UMODEL_METRICSET_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold", "90"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12920_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "content_template", "empty array repro test modified"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "UMODEL"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.0", "action-003"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.0.enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_type", "instance"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "UMODEL_METRICSET_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_domain", "ecs"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric", "CPUUtilization"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric_set", "acs_ecs_dashboard"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "empty-array-repro-modify-nonempty"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "UMODEL_METRICSET_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold", "95"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12920_step2(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "content_template", "empty array repro test emptied"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "UMODEL"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.#", "0"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.0.enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_type", "instance"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "UMODEL_METRICSET_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_domain", "ecs"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric", "CPUUtilization"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric_set", "acs_ecs_dashboard"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "empty-array-repro-modify-empty-array-repro"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "UMODEL_METRICSET_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold", "95"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12920_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "empty array repro test"

        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "UMODEL"
    }

        action_integration_config {
        actions = ["action-001", "action-002"]
        enabled = true
    }

        arms_integration_config {
        enabled = false
    }

        query_config {
        entity_type = "instance"
        type = "UMODEL_METRICSET_QUERY"
        entity_domain = "ecs"
        metric = "CPUUtilization"
        metric_set = "acs_ecs_dashboard"
    }

        display_name = "empty-array-repro-init"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "UMODEL_METRICSET_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
        threshold = 90
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12920_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "empty array repro test modified"

        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "UMODEL"
    }

        action_integration_config {
        actions = ["action-003"]
        enabled = true
    }

        arms_integration_config {
        enabled = false
    }

        query_config {
        entity_type = "instance"
        type = "UMODEL_METRICSET_QUERY"
        entity_domain = "ecs"
        metric = "CPUUtilization"
        metric_set = "acs_ecs_dashboard"
    }

        display_name = "empty-array-repro-modify-nonempty"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "UMODEL_METRICSET_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
        threshold = 95
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12920_step2(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "empty array repro test emptied"

        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "UMODEL"
    }

        action_integration_config {
        actions = []
        enabled = true
    }

        arms_integration_config {
        enabled = false
    }

        query_config {
        entity_type = "instance"
        type = "UMODEL_METRICSET_QUERY"
        entity_domain = "ecs"
        metric = "CPUUtilization"
        metric_set = "acs_ecs_dashboard"
    }

        display_name = "empty-array-repro-modify-empty-array-repro"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "UMODEL_METRICSET_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
        threshold = 95
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_APM_YOY_test 12918
func TestAccAliCloudCmsAlertRuleV2_basic12918(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12918_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.status.count"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.1", "status"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "status"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-svc-yoy-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-svc-yoy-2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-yoy"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "YOY_UP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "AVG"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.yoy_time_value", "5"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "30"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.yoy_time_unit", "minute"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12918_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.status.count"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.1", "status"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "status"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-svc-yoy-3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-svc-yoy-4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.2", "arms-svc-yoy-5"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-yoy-week"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "YOY_DOWN"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "SUM"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.yoy_time_value", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "20"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.yoy_time_unit", "week"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12918_step2(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.status.count"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.0", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.group_by.1", "status"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.window_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.0.key", "rpc"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.type", "ALL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.filter_list.1.key", "status"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "arms-svc-yoy-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.1", "arms-svc-yoy-2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-yoy-month"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "YOY_UP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "APM_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.aggregate", "MAX"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.yoy_time_value", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.severity", "ERROR"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold_list.0.threshold", "50"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.yoy_time_unit", "month"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12918_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.status.count"
            group_by = ["rpc", "status"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        filter_list {
            type = "ALL"
            key = "status"
        }
        service_id_list = ["arms-svc-yoy-1", "arms-svc-yoy-2"]
    }

        display_name = "regression-apm-yoy"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "YOY_UP"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "AVG"
        yoy_time_value = 5
        threshold_list {
            severity = "CRITICAL"
            threshold = 30
        }
        yoy_time_unit = "minute"
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12918_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.status.count"
            group_by = ["rpc", "status"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        filter_list {
            type = "ALL"
            key = "status"
        }
        service_id_list = ["arms-svc-yoy-3", "arms-svc-yoy-4", "arms-svc-yoy-5"]
    }

        display_name = "regression-apm-yoy-week"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "YOY_DOWN"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "SUM"
        yoy_time_value = 2
        threshold_list {
            severity = "WARNING"
            threshold = 20
        }
        yoy_time_unit = "week"
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12918_step2(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "APM"
    }

        query_config {
        measure_list {
            measure_code = "appstat.status.count"
            group_by = ["rpc", "status"]
            window_secs = 60
        }
        type = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key = "rpc"
        }
        filter_list {
            type = "ALL"
            key = "status"
        }
        service_id_list = ["arms-svc-yoy-1", "arms-svc-yoy-2"]
    }

        display_name = "regression-apm-yoy-month"

        enabled = false

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "YOY_UP"
        type = "APM_SIMPLE_CONDITION"
        aggregate = "MAX"
        yoy_time_value = 1
        threshold_list {
            severity = "ERROR"
            threshold = 50
        }
        yoy_time_unit = "month"
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_ExpressEscalation_test 12931
func TestAccAliCloudCmsAlertRuleV2_basic12931(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12931_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes < 0.2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-prom-express-escalation"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12931_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes < 0.1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-prom-express-escalation-updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12931_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes < 0.2"
    }

        display_name = "regression-prom-express-escalation"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12931_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes < 0.1"
    }

        display_name = "regression-prom-express-escalation-updated"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_Prometheus_Basic_test 12933
func TestAccAliCloudCmsAlertRuleV2_basic12933(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12933_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-basic-prometheus"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12933_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-basic-prometheus-updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test-updated"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12933_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 1"
    }

        display_name = "regression-basic-prometheus"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12933_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
    }

        display_name = "regression-basic-prometheus-updated"

        enabled = false

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test-updated"]
        }
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_lifecycle_test 12927
func TestAccAliCloudCmsAlertRuleV2_basic12927(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12927_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "content_template"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "UMODEL"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.0.enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_type", "instance"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "UMODEL_METRICSET_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_domain", "ecs"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric", "CPUUtilization"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric_set", "acs_ecs_dashboard"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "cspec-lifecycle-test-rule"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "UMODEL_METRICSET_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold", "90"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12927_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "content_template"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "UMODEL"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.actions.0", "action-integration-001"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.0.enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_type", "instance"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "UMODEL_METRICSET_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_domain", "ecs"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric", "CPUUtilization"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric_set", "acs_ecs_dashboard"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "cspec-lifecycle-test-rule-updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "UMODEL_METRICSET_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold", "95"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12927_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "alert triggered on $${metric}"

        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "UMODEL"
    }

        action_integration_config {
        enabled = false
    }

        arms_integration_config {
        enabled = false
    }

        query_config {
        entity_type = "instance"
        type = "UMODEL_METRICSET_QUERY"
        entity_domain = "ecs"
        metric = "CPUUtilization"
        metric_set = "acs_ecs_dashboard"
    }

        display_name = "cspec-lifecycle-test-rule"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "UMODEL_METRICSET_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
        threshold = 90
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12927_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "alert updated on $${metric}"

        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "UMODEL"
    }

        action_integration_config {
        actions = ["action-integration-001"]
        enabled = true
    }

        arms_integration_config {
        enabled = true
    }

        query_config {
        entity_type = "instance"
        type = "UMODEL_METRICSET_QUERY"
        entity_domain = "ecs"
        metric = "CPUUtilization"
        metric_set = "acs_ecs_dashboard"
    }

        display_name = "cspec-lifecycle-test-rule-updated"

        enabled = false

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "UMODEL_METRICSET_CONDITION"
        severity = "CRITICAL"
        duration_secs = 120
        threshold = 95
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_Prometheus_NotifyStrategies_test 12926
func TestAccAliCloudCmsAlertRuleV2_basic12926(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12926_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-strategies"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "NOTIFY_POLICY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.0", "strategy-1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12926_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-strategies-m1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "NOTIFY_POLICY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.0", "strategy-id-003"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12926_step2(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-strategies-m2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "NOTIFY_POLICY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.0", "strategy-4"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12926_step3(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-strategies-m3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "NOTIFY_POLICY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.0", "strategy-5"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12926_step4(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.region_id", "cn-hangzhou"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-notify-strategies-m4"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "NOTIFY_POLICY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.notify_strategies.#", "0"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12926_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 1"
    }

        display_name = "regression-notify-strategies"

        enabled = true

        notify_config {
        type = "NOTIFY_POLICY"
        notify_strategies = ["strategy-1"]
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12926_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
    }

        display_name = "regression-notify-strategies-m1"

        enabled = true

        notify_config {
        type = "NOTIFY_POLICY"
        notify_strategies = ["strategy-id-003"]
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12926_step2(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
    }

        display_name = "regression-notify-strategies-m2"

        enabled = true

        notify_config {
        type = "NOTIFY_POLICY"
        notify_strategies = ["strategy-4"]
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12926_step3(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
    }

        display_name = "regression-notify-strategies-m3"

        enabled = true

        notify_config {
        type = "NOTIFY_POLICY"
        notify_strategies = ["strategy-5"]
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12926_step4(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
        region_id = "cn-hangzhou"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
    }

        display_name = "regression-notify-strategies-m4"

        enabled = true

        notify_config {
        type = "NOTIFY_POLICY"
        notify_strategies = []
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case delete_minimal 12925
func TestAccAliCloudCmsAlertRuleV2_basic12925(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12925_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "content_template", "minimal delete test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "UMODEL"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "action_integration_config.0.enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "arms_integration_config.0.enabled", "false"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_type", "instance"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "UMODEL_METRICSET_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.entity_domain", "ecs"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric", "CPUUtilization"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.metric_set", "acs_ecs_dashboard"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "cspec-delete-minimal-test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.operator", "GT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "UMODEL_METRICSET_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.threshold", "90"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12925_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "minimal delete test"

        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "UMODEL"
    }

        action_integration_config {
        enabled = false
    }

        arms_integration_config {
        enabled = false
    }

        query_config {
        entity_type = "instance"
        type = "UMODEL_METRICSET_QUERY"
        entity_domain = "ecs"
        metric = "CPUUtilization"
        metric_set = "acs_ecs_dashboard"
    }

        display_name = "cspec-delete-minimal-test"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        operator = "GT"
        type = "UMODEL_METRICSET_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
        threshold = 90
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Case resource_AlertRuleV2_Labels_Verify_test 12938
func TestAccAliCloudCmsAlertRuleV2_basic12938(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12938_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "annotations.%", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "annotations.summary", "CPU usage high"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "annotations.description", "Node CPU is above 90%"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "cloudspec_labels_verify_test"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "labels.%", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "labels.severity", "critical"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "labels.env", "test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "CRITICAL"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "60"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("datasource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("notify_strategy_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("observe_resource_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("partition_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("severity_levels"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_cms_alert_rule_v2.default", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2_basic12938_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "120"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "PROMETHEUS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.instance_id", "prom-regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "PROMETHEUS_SINGLE_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "up > 2"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "annotations.%", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "annotations.summary", "CPU usage updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "annotations.description", "Updated CPU alert"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "annotations.runbook", "https://example.com"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "cloudspec_labels_verify_test_updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "labels.%", "3"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "labels.severity", "warning"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "labels.team", "dev"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "labels.env", "staging"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "PROMETHEUS_SIMPLE_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.duration_secs", "120"),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_basic12938_step0(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 1"
    }

        annotations = {
        "summary" = "CPU usage high"
        "description" = "Node CPU is above 90%%"
    }

        display_name = "cloudspec_labels_verify_test"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        labels = {
        "severity" = "critical"
        "env" = "test"
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "CRITICAL"
        duration_secs = 60
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

func testAccAliCloudCmsAlertRuleV2_basic12938_step1(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        schedule_config {
        type = "FIXED"
        interval_secs = 120
    }

        datasource_config {
        type = "PROMETHEUS"
        instance_id = "prom-regression-test"
    }

        query_config {
        type = "PROMETHEUS_SINGLE_QUERY"
        expr = "up > 2"
    }

        annotations = {
        "summary" = "CPU usage updated"
        "description" = "Updated CPU alert"
        "runbook" = "https://example.com"
    }

        display_name = "cloudspec_labels_verify_test_updated"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        labels = {
        "severity" = "warning"
        "team" = "dev"
        "env" = "staging"
    }

        condition_config {
        type = "PROMETHEUS_SIMPLE_CONDITION"
        severity = "WARNING"
        duration_secs = 120
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Test Cms AlertRuleV2 log based alert rules (SLS datasource and UMODEL logset query).
// Cases added by RD: the CloudSpec regression cases above do not cover the log datasource
// (project/stores) or the logset query (log_set) added for the log based alert rule request.
func TestAccAliCloudCmsAlertRuleV2_slsLogDatasource(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2SlsLogDatasource(name, "sls log alert", "WARNING", 1, "INFO"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "SLS"),
					resource.TestCheckResourceAttrPair("alicloud_cms_alert_rule_v2.default", "datasource_config.0.project", "alicloud_log_project.default", "project_name"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.stores.#", "1"),
					resource.TestCheckResourceAttrPair("alicloud_cms_alert_rule_v2.default", "datasource_config.0.stores.0.store", "alicloud_log_store.default", "logstore_name"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.stores.0.store_type", "LOG"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "SLS_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.group_type", "none"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.queries.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.queries.0.expr", "* | select count(*) as cnt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.queries.0.time_unit", "minute"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.queries.0.window", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.type", "FIXED"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "schedule_config.0.interval_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "SLS_MULTI_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.triggers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.triggers.0.operator", "HAS_DATA_MATCH_COUNT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.triggers.0.condition", "cnt > 0"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.triggers.0.count_condition", "__count__ > 0"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.triggers.0.severity", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.alert_count", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.no_data_alert_level", "INFO"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "sls log alert"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.type", "DIRECT_NOTIFY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.type", "GROUP"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),

					resource.TestCheckResourceAttrSet("alicloud_cms_alert_rule_v2.default", "observe_resource_global_scope"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2SlsLogDatasource(name, "sls log alert updated", "ERROR", 2, "WARNING"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "SLS"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.stores.#", "1"),
					resource.TestCheckResourceAttrPair("alicloud_cms_alert_rule_v2.default", "datasource_config.0.stores.0.store", "alicloud_log_store.default", "logstore_name"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "SLS_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.group_type", "none"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.queries.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.triggers.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.triggers.0.operator", "HAS_DATA_MATCH_COUNT"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.triggers.0.count_condition", "__count__ > 0"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.triggers.0.severity", "ERROR"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.alert_count", "2"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.no_data_alert_level", "WARNING"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "sls log alert updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2SlsLogDatasource(name, displayName, severity string, alertCount int, noDataAlertLevel string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
    description  = "tf-acc cms alert rule v2 sls"
}

resource "alicloud_log_store" "default" {
    provider = alicloud.hz
    project_name = alicloud_log_project.default.project_name
    logstore_name = var.name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "sls log alert rule"

        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "SLS"
        project = alicloud_log_project.default.project_name
        stores {
            store = alicloud_log_store.default.logstore_name
            store_type = "LOG"
        }
    }

        query_config {
        type = "SLS_MULTI_QUERY"
        group_type = "none"
        queries {
            expr = "* | select count(*) as cnt"
            time_unit = "minute"
            window = 60
        }
    }

        display_name = "%[2]s"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "SLS_MULTI_CONDITION"
        triggers {
            operator = "HAS_DATA_MATCH_COUNT"
            condition = "cnt > 0"
            count_condition = "__count__ > 0"
            severity = "%[3]s"
        }
        alert_count = %[4]d
        no_data_alert_level = "%[5]s"
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name, displayName, severity, alertCount, noDataAlertLevel)
}

func TestAccAliCloudCmsAlertRuleV2_umodelLogSet(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2UmodelLogSet(name, "umodel logset alert"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "UMODEL"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "UMODEL_LOGSET_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.log_set", name),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.expr", "* | select count(*) as cnt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.window_secs", "60"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.type", "UMODEL_LOGSET_CONDITION"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "condition_config.0.count_threshold", "1"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "umodel logset alert"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "true"),

					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "notify_config.0.channels.0.identifiers.0", "regression-test"),
				),
			},
			{
				Config: testAccAliCloudCmsAlertRuleV2UmodelLogSet(name, "umodel logset alert updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.log_set", name),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "umodel logset alert updated"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),
				),
			},
			{
				ResourceName:            "alicloud_cms_alert_rule_v2.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2UmodelLogSet(name, displayName string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
        provider = alicloud.hz
        content_template = "umodel logset alert rule"

        schedule_config {
        type = "FIXED"
        interval_secs = 60
    }

        datasource_config {
        type = "UMODEL"
    }

        query_config {
        type = "UMODEL_LOGSET_QUERY"
        log_set = var.name
        expr = "* | select count(*) as cnt"
        window_secs = 60
    }

        display_name = "%[2]s"

        enabled = true

        notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type = "GROUP"
            identifiers = ["regression-test"]
        }
    }

        condition_config {
        type = "UMODEL_LOGSET_CONDITION"
        severity = "CRITICAL"
        count_operator = "GT"
        count_threshold = 1
        match_field = "level"
        match_operator = "EQUALS"
        match_value = "ERROR"
    }

        workspace = alicloud_cms_workspace.default.id
}
`, name, displayName)
}

// Case resource_AlertRuleV2_ObserveResource_Rca_test validates the observe_resource_config and rca_config blocks.
func TestAccAliCloudCmsAlertRuleV2_observeResourceRca(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", r)
	resource.Test(t, resource.TestCase{
		CheckDestroy: checkCmsAlertRuleV2Destroy,
		PreCheck: func() {
			acctest.PreCheck(t)
			if factory := acctest.ProtoV5ProviderFactories()["alicloud"]; factory == nil {
				t.Fatal("SDK acceptance tests require the real alicloud protocol 5 factory")
			}
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudCmsAlertRuleV2_observeResourceRca(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "datasource_config.0.type", "APM"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.type", "APM_MULTI_QUERY"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.measure_list.0.measure_code", "appstat.transaction.rt"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "query_config.0.service_id_list.0", "svc-coverage-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "observe_resource_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "observe_resource_config.0.entity_domain", "apm"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "observe_resource_config.0.entity_type", "apm.service"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "observe_resource_config.0.relation_type", "CLOUD_INSTANCE"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "observe_resource_config.0.resources.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "observe_resource_config.0.resources.0", "svc-coverage-1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "rca_config.#", "1"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "rca_config.0.enable_rca", "true"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "rca_config.0.digital_employee_name", "tf-probe-agent"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "display_name", "regression-apm-observe-rca"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "enabled", "false"),
					resource.TestCheckResourceAttr("alicloud_cms_alert_rule_v2.default", "workspace", name),
				),
			},
			{
				ResourceName:      "alicloud_cms_alert_rule_v2.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAliCloudCmsAlertRuleV2_observeResourceRca(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

provider "alicloud" {
    alias  = "hz"
    region = "cn-hangzhou"
}

resource "alicloud_log_project" "default" {
    provider = alicloud.hz
    project_name = var.name
}

resource "alicloud_cms_workspace" "default" {
    provider = alicloud.hz
    workspace_name = var.name
    sls_project    = alicloud_log_project.default.project_name
}

resource "alicloud_cms_alert_rule_v2" "default" {
    provider = alicloud.hz
    schedule_config {
        type          = "FIXED"
        interval_secs = 60
    }
    datasource_config {
        type = "APM"
    }
    query_config {
        measure_list {
            measure_code = "appstat.transaction.rt"
            group_by     = ["rpc"]
            window_secs  = 60
        }
        type            = "APM_MULTI_QUERY"
        filter_list {
            type = "ALL"
            key  = "rpc"
        }
        service_id_list = ["svc-coverage-1"]
    }
    observe_resource_config {
        entity_domain = "apm"
        entity_type   = "apm.service"
        relation_type = "CLOUD_INSTANCE"
        resources     = ["svc-coverage-1"]
    }
    rca_config {
        enable_rca           = true
        digital_employee_name = "tf-probe-agent"
    }
    display_name = "regression-apm-observe-rca"
    enabled      = false
    notify_config {
        type = "DIRECT_NOTIFY"
        channels {
            type        = "GROUP"
            identifiers = ["regression-test"]
        }
    }
    condition_config {
        operator  = "GT"
        type      = "APM_SIMPLE_CONDITION"
        aggregate = "AVG"
        threshold_list {
            severity  = "WARNING"
            threshold = 100
        }
    }
    workspace = alicloud_cms_workspace.default.id
}
`, name)
}

// Test Cms AlertRuleV2. <<< Resource test cases, automatically generated.
