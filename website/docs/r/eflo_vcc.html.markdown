---
subcategory: "Eflo"
layout: "alicloud"
page_title: "Alicloud: alicloud_eflo_vcc"
description: |-
  Provides a Alicloud Eflo Vcc resource.
---

# alicloud_eflo_vcc

Provides a Eflo Vcc resource.



For information about Eflo Vcc and how to use it, see [What is Vcc](https://next.api.alibabacloud.com/document/eflo/2022-05-30/CreateVcc).

-> **NOTE:** Available since v2.2.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-wulanchabu"
}

variable "region_id" {
  default = "cn-wulanchabu"
}

variable "zone_id" {
  default = "cn-wulanchabu-b"
}

resource "alicloud_vpc" "VPC" {
  is_default  = false
  description = "vccexample用"
  cidr_block  = "192.168.0.0/16"
  enable_ipv6 = false
  vpc_name    = "vpc-vcc-example-tf"
}

resource "alicloud_vswitch" "VSW" {
  zone_id      = var.zone_id
  cidr_block   = "192.168.0.0/24"
  vswitch_name = "vsw-vcc-tf"
  vpc_id       = alicloud_vpc.VPC.id
}

resource "alicloud_cen_instance" "CEN" {
  cen_instance_name = "cen-tf-example-1"
}

resource "alicloud_cen_transit_router" "CEN-TR" {
  cen_id              = alicloud_cen_instance.CEN.id
  transit_router_name = "tf-example"
}

resource "alicloud_cen_transit_router_vpc_attachment" "VPC-ATTACH" {
  vpc_id                             = alicloud_vpc.VPC.id
  transit_router_vpc_attachment_name = "tf-example"
  cen_id                             = alicloud_cen_instance.CEN.id
  zone_mappings {
    vswitch_id = alicloud_vswitch.VSW.id
    zone_id    = var.zone_id
  }
  transit_router_id = alicloud_cen_transit_router.CEN-TR.transit_router_id
}

resource "alicloud_vpc" "VPC-CRAETE" {
  is_default  = false
  description = "TFexample"
  cidr_block  = "192.168.0.0/16"
  enable_ipv6 = false
  vpc_name    = "vpc-vcc-example-tf-2"
}

resource "alicloud_vswitch" "VSW-CREATE" {
  vpc_id       = alicloud_vpc.VPC-CRAETE.id
  zone_id      = var.zone_id
  cidr_block   = "192.168.0.0/24"
  vswitch_name = "vsw-vcc-tf-2"
}


resource "alicloud_eflo_vcc" "default" {
  description          = "example-tf"
  access_could_service = false
  connection_type      = "VPC"
  zone_id              = var.zone_id
  vswitch_id           = alicloud_vswitch.VSW-CREATE.id
  vpc_id               = alicloud_vpc.VPC-CRAETE.id
  bgp_asn              = 64512
  bgp_cidr             = "169.254.128.0/23"
  bandwidth            = 1000
  vcc_name             = "vcc-tf"
}
```

## Argument Reference

The following arguments are supported:
* `access_could_service` - (Optional) Start access to cloud service, optional value:
  - `true`: enable access to cloud services
  - `false`: does not enable access to cloud services

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `bandwidth` - (Required, Int) Bandwidth, in Mbps. The minimum value is 1000, representing 1Gbps bandwidth; The maximum value is 400000, representing 400Gbps bandwidth.

-> **NOTE:**  1Gbps = 1000Mbps

* `bgp_asn` - (Optional, Int) Bgp as number

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `bgp_cidr` - (Optional, ForceNew, Computed) BGP network segment.
* `cen_id` - (Optional) CEN
* `cen_owner_id` - (Optional, ForceNew) Account to which cen belongs
* `connection_type` - (Required) Connection method, optional value:
  - `VPC`
  - `CENTR`
* `description` - (Optional) Description

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `resource_group_id` - (Optional, Computed) Resource group id
* `tags` - (Optional, Map) The tag of the resource
* `vswitch_id` - (Required) vSwitch
* `vcc_name` - (Optional) Lingjun Connection Name
* `vpc_id` - (Required) VPC
* `zone_id` - (Required, ForceNew) Availability Zone

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `create_time` - Creation time.
* `gmt_modified` - Modification time.
* `message` - Information returned when the call fails.
* `spec` - Specifications.
* `status` - Status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 22 mins) Used when create the Vcc.
* `delete` - (Defaults to 62 mins) Used when delete the Vcc.
* `update` - (Defaults to 62 mins) Used when update the Vcc.

## Import

Eflo Vcc can be imported using the id, e.g.

```shell
$ terraform import alicloud_eflo_vcc.example <vcc_id>
```