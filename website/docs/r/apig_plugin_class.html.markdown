---
subcategory: "Cloud Native API Gateway (APIG)"
layout: "alicloud"
page_title: "Alicloud: alicloud_apig_plugin_class"
description: |-
  Provides a Alicloud APIG Plugin Class resource.
---

# alicloud_apig_plugin_class

Provides a APIG Plugin Class resource.

plugin class info.

For information about APIG Plugin Class and how to use it, see [What is Plugin Class](https://next.api.alibabacloud.com/document/APIG/2024-03-27/CreatePluginClass).

-> **NOTE:** Available since v1.285.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_apig_plugin_class" "default" {
  wasm_url                      = "https://example.com/plugin.wasm"
  description                   = "A test plugin class"
  version_description           = "Initial version"
  plugin_class_name             = "example-plugin-class"
  version                       = "1.0.2"
  alias                         = "example-plugin-alias"
  execute_priority              = 1
  wasm_language                 = "TinyGo"
  supported_min_gateway_version = "1.0.0"
  execute_stage                 = "UNSPECIFIED_PHASE"
}
```

## Argument Reference

The following arguments are supported:
* `alias` - (Optional, ForceNew) plugin class alias
* `description` - (Required, ForceNew) plugin class description 
* `execute_priority` - (Required, Int, ForceNew) Plugin execute priority. Minimum value is 1.
* `execute_stage` - (Required, ForceNew) Plugin execute stage. Valid values: `AUTHN`, `AUTHZ`, `STATS`, `UNSPECIFIED_PHASE`.
* `plugin_class_name` - (Required, ForceNew) The name of the plugin class.
* `supported_min_gateway_version` - (Optional, ForceNew) The minimum supported gateway version.
* `version` - (Required, ForceNew) The plugin class version.
* `version_description` - (Required, ForceNew) The version description.
* `wasm_language` - (Required, ForceNew) The WASM language. Valid values: `TinyGo`, `Cpp`, `Rust`, `AssemblyScript`, `Zig`.
* `wasm_url` - (Required, ForceNew) The URL of the WASM file.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `document` - plugin document.
* `status` - publish status of plugin class.
* `type` - plugin class type.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 10 mins) Used when create the Plugin Class.
* `delete` - (Defaults to 5 mins) Used when delete the Plugin Class.

## Import

APIG Plugin Class can be imported using the id, e.g.

```shell
$ terraform import alicloud_apig_plugin_class.example <plugin_class_id>
```