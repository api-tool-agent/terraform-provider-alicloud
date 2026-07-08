---
subcategory: "Cloud Native API Gateway (APIG)"
layout: "alicloud"
page_title: "Alicloud: alicloud_apig_plugin_classes"
description: |-
  Provides a list of APIG Plugin Classes in Alicloud.
---

# alicloud_apig_plugin_classes

This data source provides the list of APIG Plugin Classes.

-> **NOTE:** Available since v1.285.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_apig_plugin_classes" "example" {}

output "first_plugin_class_id" {
  value = data.alicloud_apig_plugin_classes.example.plugin_classes[0].id
}
```

## Argument Reference

The following arguments are supported:

* `ids` - (Optional) A list of Plugin Class IDs to filter the results.
* `name_regex` - (Optional) A regex string to filter results by plugin class name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).
* `type` - (Optional) Filter by plugin class type. Valid values: `Auth`, `FlowControl`, `FlowObservation`, `Security`, `TransportProtocol`, `Other`.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:

* `ids` - A list of Plugin Class IDs.
* `names` - A list of plugin class names.
* `plugin_classes` - A list of Plugin Classes. Each element contains the following attributes:
  * `alias` - The alias of the plugin class.
  * `description` - The description of the plugin class.
  * `document` - The plugin class document.
  * `id` - The ID of the plugin class.
  * `plugin_class_id` - The plugin class ID.
  * `plugin_class_name` - The name of the plugin class.
  * `status` - The publish status of the plugin class.
  * `type` - The type of the plugin class.
  * `version` - The version of the plugin class.
  * `wasm_language` - The WASM language of the plugin class.
