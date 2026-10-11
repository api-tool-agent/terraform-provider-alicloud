---
subcategory: "Pai Lang Studio"
layout: "alicloud"
page_title: "Alicloud: alicloud_pai_lang_studio_toolsets"
sidebar_current: "docs-alicloud-datasource-pai-lang-studio-toolsets"
description: |-
  Provides a list of Pai Lang Studio Toolset owned by an Alibaba Cloud account.
---

# alicloud_pai_lang_studio_toolsets

This data source provides Pai Lang Studio Toolset available to the user.[What is Toolset](https://next.api.alibabacloud.com/document/PAILangStudio/2024-07-10/CreateToolset)

-> **NOTE:** Available since v1.295.0.

## Example Usage

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

data "alicloud_pai_lang_studio_toolsets" "default" {
  workspace_id = alicloud_pai_workspace_workspace.default.id
  ids          = [alicloud_pai_lang_studio_toolset.default.id]
  name_regex   = alicloud_pai_lang_studio_toolset.default.toolset_name
}

output "toolset_id" {
  value = data.alicloud_pai_lang_studio_toolsets.default.toolsets.0.toolset_id
}
```

## Argument Reference

The following arguments are supported:
* `accessibility` - (Optional) The visibility of the toolset. Valid values: `PUBLIC`, `PRIVATE`.
* `creator` - (Optional) The creator of the toolset.
* `order` - (Optional) The sort order of the results.
* `sort_by` - (Optional) The field used to sort the results.
* `toolset_id` - (Optional) The ID of the toolset.
* `toolset_name` - (Optional) The name of the toolset.
* `toolset_type` - (Optional) The type of the toolset. Valid values: `OPENAPI`.
* `workspace_id` - (Optional) The ID of the workspace to which the toolset belongs.
* `ids` - (Optional, Computed) A list of Toolset IDs. The value is formulated as `<workspace_id>:<toolset_id>`.
* `name_regex` - (Optional) A regex string to filter results by toolset name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).


## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Toolset IDs.
* `names` - A list of name of Toolsets.
* `toolsets` - A list of Toolset Entries. Each element contains the following attributes:
    * `accessibility` - The visibility of the toolset.
    * `creator` - The creator of the toolset.
    * `description` - The description of the toolset.
    * `gmt_create_time` - The creation time of the toolset.
    * `gmt_modified_time` - The modification time of the toolset.
    * `tools` - The list of tools contained in the toolset.
    * `toolset_config` - The configuration of the toolset.
    * `toolset_id` - The ID of the toolset.
    * `toolset_name` - The name of the toolset.
    * `toolset_type` - The type of the toolset.
    * `workspace_id` - The ID of the workspace to which the toolset belongs.
    * `id` - The ID of the resource supplied above.
