---
subcategory: "PAI Workspace"
layout: "alicloud"
page_title: "Alicloud: alicloud_pai_workspace_config"
description: |-
  Provides a Alicloud PAI Workspace Cloud Config (Config) resource.
---

# alicloud_pai_workspace_config

Provides a PAI Workspace Cloud Config (Config) resource.



For information about PAI Workspace Cloud Config (Config) and how to use it, see [What is Cloud Config (Config)](https://next.api.alibabacloud.com/document/AIWorkSpace/2021-02-04/UpdateConfig).

-> **NOTE:** Available since v1.287.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_pai_workspace_workspace" "workspace" {
  description    = "函数引用： random(2,100)"
  display_name   = "example_pop_config_25"
  workspace_name = "example_pop_config_982"
  env_types      = ["prod"]
}


resource "alicloud_pai_workspace_config" "default" {
  config_value = "\"[{\\\"Type\\\":\\\"Quota\\\",\\\"Key\\\":\\\"quota16a5gjqy2np\\\",\\\"Value\\\":[{\\\"Products\\\":\\\"DLC,DSW\\\",\\\"SpecName\\\":\\\"aaa\\\",\\\"Spec\\\":{\\\"CPU\\\":\\\"8\\\",\\\"Memory\\\":\\\"4Gi\\\",\\\"GPU\\\":\\\"1\\\"}}]},{\\\"Type\\\":\\\"Quota\\\",\\\"Key\\\":\\\"quotabn1xxh0uc7j\\\",\\\"Value\\\":[{\\\"SpecName\\\":\\\"examplelinglan\\\",\\\"Spec\\\":{\\\"CPU\\\":\\\"1\\\",\\\"GPU\\\":\\\"1\\\",\\\"Memory\\\":\\\"1Gi\\\"},\\\"Products\\\":\\\"DSW\\\"}]}]\""
  workspace_id = alicloud_pai_workspace_workspace.workspace.id
  labels {
    key   = "system.categoryName1"
    value = "DSWAutoRecycle"
  }
  labels {
    key   = "system.categoryName2"
    value = "DSWAutoRecycle"
  }
  labels {
    key   = "system.categoryName3"
    value = "DSWAutoRecycle"
  }
  config_key    = "resourceSpecs"
  category_name = "CommonQuotaConfig"
}
```

### Deleting `alicloud_pai_workspace_config` or removing it from your configuration

Terraform cannot destroy resource `alicloud_pai_workspace_config`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `category_name` - (Optional, ForceNew, Computed) Resource attribute field representing classification name
* `config_key` - (Optional, ForceNew, Computed) Resource attribute field representing the workspace configuration name
* `config_value` - (Optional) Resource labels representing workspace configuration values
* `labels` - (Optional, List) Representative Workspace Configuration Label See [`labels`](#labels) below.
* `workspace_id` - (Required, ForceNew) Resource attribute field representing workspace ID

### `labels`

The labels supports the following:
* `key` - (Optional) Resource representing the tag Key
* `value` - (Optional) Resource representing the label Value

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<workspace_id>:<category_name>:<config_key>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Cloud Config (Config).
* `update` - (Defaults to 5 mins) Used when update the Cloud Config (Config).

## Import

PAI Workspace Cloud Config (Config) can be imported using the id, e.g.

```shell
$ terraform import alicloud_pai_workspace_config.example <workspace_id>:<category_name>:<config_key>
```