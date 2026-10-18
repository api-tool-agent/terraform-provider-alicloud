---
subcategory: "PaiStudio"
layout: "alicloud"
page_title: "Alicloud: alicloud_paistudio_algorithm"
description: |-
  Provides a Alicloud PaiStudio Algorithm resource.
---

# alicloud_paistudio_algorithm

Provides a PaiStudio Algorithm resource.

For information about PaiStudio Algorithm and how to use it, see [What is Algorithm](https://next.api.alibabacloud.com/document/PaiStudio/2022-01-12/CreateAlgorithm).

-> **NOTE:** Available since v2.1.0.

## Example Usage

Basic Usage

```terraform
resource "alicloud_paistudio_algorithm" "example" {
  algorithm_name        = "terraform-example-algorithm"
  display_name          = "terraform-example-algorithm-display"
  algorithm_description = "Created by Terraform"
}
```

## Argument Reference

The following arguments are supported:
* `algorithm_name` - (Required, ForceNew, String) The name of the Algorithm. The UpdateAlgorithm API only accepts the description and display name, so changing the name recreates the Algorithm.
* `algorithm_description` - (Optional, Computed, String) The description of the Algorithm.
* `display_name` - (Optional, Computed, String) The display name of the Algorithm.
* `workspace_id` - (Optional, ForceNew, Computed, String) The ID of the workspace that the Algorithm belongs to. If not set, the API assigns the default workspace; changing it recreates the Algorithm.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `algorithm_provider` - The provider of the Algorithm.
* `gmt_create_time` - The time when the Algorithm was created.
* `gmt_modified_time` - The time when the Algorithm was last modified.
* `tenant_id` - The ID of the tenant that the Algorithm belongs to.
* `user_id` - The ID of the user who owns the Algorithm.

## Import

PaiStudio Algorithm can be imported using the id, e.g.

```shell
$ terraform import alicloud_paistudio_algorithm.example <algorithm_id>
```
