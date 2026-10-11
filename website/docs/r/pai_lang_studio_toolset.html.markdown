---
subcategory: "Pai Lang Studio"
layout: "alicloud"
page_title: "Alicloud: alicloud_pai_lang_studio_toolset"
description: |-
  Provides a Alicloud Pai Lang Studio Toolset resource.
---

# alicloud_pai_lang_studio_toolset

Provides a Pai Lang Studio Toolset resource.

For information about Pai Lang Studio Toolset and how to use it, see [What is Toolset](https://next.api.alibabacloud.com/document/PAILangStudio/2024-07-10/CreateToolset).

-> **NOTE:** Available since v1.295.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

resource "alicloud_pai_workspace_workspace" "default" {
  description    = "for toolset example"
  display_name   = var.name
  workspace_name = var.name
  env_types      = ["prod"]
}

resource "alicloud_pai_lang_studio_toolset" "default" {
  workspace_id   = alicloud_pai_workspace_workspace.default.id
  toolset_name   = var.name
  toolset_type   = "OPENAPI"
  accessibility  = "PRIVATE"
  description    = "toolset created by terraform"
  toolset_config = jsonencode({
    openapi = "3.0.0"
    info = {
      title   = "example"
      version = "1.0.0"
    }
    paths = {
      "/example" = {
        get = {
          operationId = "example"
          responses = {
            "200" = {
              description = "ok"
            }
          }
        }
      }
    }
  })
}
```

## Argument Reference

The following arguments are supported:
* `accessibility` - (Optional, ForceNew) The visibility of the toolset. Valid values: `PUBLIC`, `PRIVATE`.
* `description` - (Optional) The description of the toolset.
* `toolset_config` - (Optional) The configuration of the toolset, in JSON string format.
* `toolset_name` - (Optional) The name of the toolset.
* `toolset_type` - (Optional, ForceNew) The type of the toolset. Valid values: `OPENAPI`.
* `workspace_id` - (Optional, ForceNew, Computed) The ID of the workspace to which the toolset belongs.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<workspace_id>:<toolset_id>`.
* `creator` - The creator of the toolset.
* `gmt_create_time` - The creation time of the toolset.
* `gmt_modified_time` - The modification time of the toolset.
* `tools` - The list of tools contained in the toolset.
* `toolset_id` - The ID of the toolset.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Toolset.
* `delete` - (Defaults to 5 mins) Used when delete the Toolset.
* `update` - (Defaults to 5 mins) Used when update the Toolset.

## Import

Pai Lang Studio Toolset can be imported using the id, e.g.

```shell
$ terraform import alicloud_pai_lang_studio_toolset.example <workspace_id>:<toolset_id>
```
