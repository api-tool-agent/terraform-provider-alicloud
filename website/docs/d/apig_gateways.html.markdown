---
subcategory: "Cloud Native API Gateway (APIG)"
layout: "alicloud"
page_title: "Alicloud: alicloud_apig_gateways"
sidebar_current: "docs-alicloud-datasource-apig-gateways"
description: |-
  Provides a list of Apig Gateway owned by an Alibaba Cloud account.
---

# alicloud_apig_gateways

This data source provides Apig Gateway available to the user.[What is Gateway](https://next.api.alibabacloud.com/document/APIG/2024-03-27/CreateGateway)

-> **NOTE:** Available since v1.285.0.

## Example Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

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
  vpc {
    vpc_id = alicloud_vpc.gateway_vpc_pre3.id
  }
  gateway_edition = "Professional"
  gateway_type    = "API"
  zones {
    zone_id    = "cn-hangzhou-i"
    vswitch_id = alicloud_vswitch.gateway_vswitch_pre3.id
  }
  vswitch {
    vswitch_id = alicloud_vswitch.gateway_vswitch_pre3.id
  }
  zone_config {
    select_option = "Auto"
  }
  payment_type = "PayAsYouGo"
  gateway_name = "gateway-alt2-test"
  spec         = "apigw.small.x1"
}

data "alicloud_apig_gateways" "default" {
  ids               = ["${alicloud_apig_gateway.default.id}"]
  name_regex        = alicloud_apig_gateway.default.gateway_name
  gateway_name      = "gateway-alt2-test"
  resource_group_id = ""
}

output "alicloud_apig_gateway_example_id" {
  value = data.alicloud_apig_gateways.default.gateways.0.id
}
```

## Argument Reference

The following arguments are supported:
* `gateway_id` - (ForceNew, Optional) Cloud-native API gateway ID.
* `gateway_name` - (ForceNew, Optional) Query by exact match of the gateway name.
* `resource_group_id` - (ForceNew, Optional) The ID of the destination resource group.
* `tags` - (Computed) The tag of the resource
* `ids` - (Optional, ForceNew, Computed) A list of Gateway IDs. 
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Group Metric Rule name.
* `enable_details` - (Optional, ForceNew) Default to `false`. Set it to `true` can output more details about resource attributes.
* `output_file` - (Optional, ForceNew) File name where to save data source results (after running `terraform plan`).


## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Gateway IDs.
* `names` - A list of name of Gateways.
* `gateways` - A list of Gateway Entries. Each element contains the following attributes:
    * `create_from` - The source from which the gateway was created.
    * `create_time` - Creation timestamp.
    * `environments` - **NOTE:** This field is only available when `enable_details` is `true`. The list of environments associated with the gateway.
        * `alias` - The alias of the environment.
        * `environment_id` - The ID of the environment.
        * `name` - The name of the environment.
    * `expire_time` - Timestamp indicating when the subscription expires.
    * `gateway_edition` - Gateway instance edition:.
    * `gateway_id` - Cloud-native API gateway ID.
    * `gateway_name` - Query by exact match of the gateway name.
    * `gateway_type` - The gateway type.
    * `load_balancers` - The list of Gateway ingress addresses.
        * `address` - The address of the load balancer for the gateway.
        * `address_ip_version` - IP version:.
        * `address_type` - Load balancer address type:.
        * `gateway_default` - Indicates whether this is the default ingress address of the gateway.
        * `ipv4_addresses` - The list of IPv4 addresses.
        * `ipv6_addresses` - The list of IPv6 addresses.
        * `load_balancer_id` - The ID of the load balancer associated with the gateway.
        * `mode` - Load balancing provisioning mode for the gateway:.
        * `ports` - The list of listening ports.
            * `port` - The port number of the load balancer listener.
            * `protocol` - Protocol:.
        * `status` - Load balancer status:.
        * `type` - Load balancer type:.
    * `payment_type` - Payment type:.
    * `resource_group_id` - The ID of the destination resource group.
    * `security_group` - Security group of the gateway.
        * `name` - **NOTE:** This field is only available when `enable_details` is `true`. The name of the security group.
        * `security_group_id` - The ID of the security group.
    * `spec` - Gateway specification:.
    * `status` - Gateway status:.
    * `sub_domain_infos` - List of second-level domain names.
        * `domain_id` - The ID of the secondary domain name for the gateway.
        * `name` - The name of the secondary domain associated with the gateway.
        * `network_type` - Network type:.
        * `protocol` - The protocol used by the secondary domain name.
    * `tags` - The tag of the resource.
    * `target_version` - The target version of the gateway instance.
    * `update_time` - The timestamp when the resource was last updated.
    * `vswitch` - **NOTE:** This field is only available when `enable_details` is `true`. The vSwitch associated with the gateway.
        * `name` - The name of the virtual switch associated with the gateway.
        * `vswitch_id` - The ID of the virtual switch.
    * `version` - The current running version of the gateway instance.
    * `vpc` - The Virtual Private Cloud (VPC) associated with the gateway.
        * `name` - **NOTE:** This field is only available when `enable_details` is `true`. The name of the VPC gateway.
        * `vpc_id` - The ID of the VPC network associated with the gateway.
    * `zones` - The list of zones associated with the gateway.
        * `name` - **NOTE:** This field is only available when `enable_details` is `true`. The name of the availability zone for the gateway.
        * `vswitch_id` - The ID of the virtual switch in the availability zone.
        * `zone_id` - The ID of the availability zone for the gateway.
    * `id` - The ID of the resource supplied above.
