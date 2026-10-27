---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_rds_parameter_group"
sidebar_current: "docs-alicloud-resource-rds-parameter-group"
description: |-
  Provides a Alicloud RDS Parameter Group resource.
---

# alicloud_rds_parameter_group

Provides a RDS Parameter Group resource.

For information about RDS Parameter Group and how to use it, see [What is Parameter Group](https://www.alibabacloud.com/help/en/doc-detail/144839.htm).

-> **NOTE:** Available since v1.119.0.

-> **NOTE:** `resource_group_id` and `modify_mode` are write-only for the RDS API: they are sent to `CreateParameterGroup`/`ModifyParameterGroup` but never returned by `DescribeParameterGroup` or `DescribeParameterGroups`. They cannot be read back, so they are empty in the state of an imported parameter group; set them in the configuration to let Terraform manage them. Removing them from the configuration keeps the current value on the server and does not fall back to any default.

## Example Usage

Basic Usage

<div style="display: block;margin-bottom: 40px;"><div class="oics-button" style="float: right;position: absolute;margin-bottom: 10px;">
  <a href="https://api.aliyun.com/terraform?resource=alicloud_rds_parameter_group&exampleId=ecba00a3-4705-c617-8963-76bd9ed57d6d16e9a7fb&activeTab=example&spm=docs.r.rds_parameter_group.0.ecba00a347&intl_lang=EN_US" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; max-width: 100%;">
  </a>
</div></div>

```terraform
variable "name" {
  default = "tf_example"
}

resource "alicloud_rds_parameter_group" "default" {
  engine         = "mysql"
  engine_version = "5.7"
  param_detail {
    param_name  = "back_log"
    param_value = "4000"
  }
  param_detail {
    param_name  = "wait_timeout"
    param_value = "86460"
  }
  parameter_group_desc = var.name
  parameter_group_name = var.name
}
```

📚 Need more examples? [VIEW MORE EXAMPLES](https://api.aliyun.com/terraform?activeTab=sample&source=Sample&sourcePath=OfficialSample:alicloud_rds_parameter_group&spm=docs.r.rds_parameter_group.example&intl_lang=EN_US)

## Argument Reference

The following arguments are supported:

* `engine` - (Required, ForceNew) The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
* `engine_version` - (Required, ForceNew) The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
* `modify_mode` - (Optional, String, Available since v2.2.0) The modification mode of the parameter template when `param_detail` is updated. Valid values: `Collectivity`, `Individual`. `Collectivity` adds or changes the parameters in `param_detail`, other parameters in the template are not affected. `Individual` replaces the parameters in the template with `param_detail`. This argument is only sent to the `ModifyParameterGroup` API when its value changes; if it is not sent, the API applies `Collectivity`. It is not returned by the API.
* `param_detail` - (Required) Parameter list. See [`param_detail`](#param_detail) below.
* `parameter_group_desc` - (Optional) The description of the parameter template.
* `parameter_group_name` - (Required) The name of the parameter template.
* `resource_group_id` - (Optional, Computed, String, Available since v2.2.0) The ID of the resource group to which the parameter template belongs. It is sent to the `CreateParameterGroup` and `ModifyParameterGroup` API. The RDS API does not return this field, so it cannot be read back after import.

### `param_detail`

The param_detail supports the following: 

* `param_name` - (Required) The name of a parameter.
* `param_value` - (Required) The value of a parameter.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID in terraform of Parameter Group.
* `related_custins_info` - (Computed, List, Available since v2.2.0) A list of RDS instances to which the parameter template is applied. Only PostgreSQL templates support this attribute. See [`related_custins_info`](#related_custins_info) below.

### `related_custins_info`

The related_custins_info supports the following:

* `db_instance_name` - (Computed, String, Available since v2.2.0) The ID of the RDS instance to which the parameter template is applied.
* `applied_time` - (Computed, String, Available since v2.2.0) The time when the parameter template was applied to the RDS instance.

## Import

RDS Parameter Group can be imported using the id, e.g.

```shell
$ terraform import alicloud_rds_parameter_group.example <id>
```
