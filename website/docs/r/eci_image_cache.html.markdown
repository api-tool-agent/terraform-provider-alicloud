---
subcategory: "Elastic Container Instance (ECI)"
layout: "alicloud"
page_title: "Alicloud: alicloud_eci_image_cache"
sidebar_current: "docs-alicloud-eci-image-cache"
description: |-
  Provides an Alicloud ECI Image Cache resource.
---

# alicloud_eci_image_cache

An ECI Image Cache can help user to solve the time-consuming problem of image pull. For information about Alicloud ECI Image Cache and how to use it, see [What is Resource Alicloud ECI Image Cache](https://www.alibabacloud.com/help/doc-detail/146891.htm).

-> **NOTE:** Available since v1.89.0.

-> **NOTE:** Each image cache corresponds to a snapshot, and the user does not delete the snapshot directly, otherwise the cache will fail.

## Example Usage

Basic Usage

<div style="display: block;margin-bottom: 40px;"><div class="oics-button" style="float: right;position: absolute;margin-bottom: 10px;">
  <a href="https://api.aliyun.com/terraform?resource=alicloud_eci_image_cache&exampleId=4f400e0a-d6a3-4598-1420-1d7f846d2e95da5df9cd&activeTab=example&spm=docs.r.eci_image_cache.0.4f400e0ad6&intl_lang=EN_US" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; max-width: 100%;">
  </a>
</div></div>

```terraform
variable "name" {
  default = "tf-example"
}

data "alicloud_eci_zones" "default" {}
resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "10.0.0.0/8"
}
resource "alicloud_vswitch" "default" {
  vswitch_name = var.name
  cidr_block   = "10.1.0.0/16"
  vpc_id       = alicloud_vpc.default.id
  zone_id      = data.alicloud_eci_zones.default.zones.0.zone_ids.0
}

resource "alicloud_security_group" "default" {
  name   = var.name
  vpc_id = alicloud_vpc.default.id
}

resource "alicloud_eip_address" "default" {
  isp                       = "BGP"
  address_name              = var.name
  netmode                   = "public"
  bandwidth                 = "1"
  security_protection_types = ["AntiDDoS_Enhanced"]
  payment_type              = "PayAsYouGo"
}
data "alicloud_regions" "default" {
  current = true
}

resource "alicloud_eci_image_cache" "default" {
  image_cache_name  = var.name
  images            = ["registry-vpc.${data.alicloud_regions.default.regions.0.id}.aliyuncs.com/eci_open/nginx:alpine"]
  security_group_id = alicloud_security_group.default.id
  vswitch_id        = alicloud_vswitch.default.id
  eip_instance_id   = alicloud_eip_address.default.id
}

```

📚 Need more examples? [VIEW MORE EXAMPLES](https://api.aliyun.com/terraform?activeTab=sample&source=Sample&sourcePath=OfficialSample:alicloud_eci_image_cache&spm=docs.r.eci_image_cache.example&intl_lang=EN_US)

## Argument Reference

The following arguments are supported:

* `image_cache_name` - (Required, ForceNew) The name of the image cache.
* `images` - (Required, ForceNew) The images to be cached. The image name must be versioned.
* `security_group_id` - (Required, ForceNew) The ID of the security group. You do not need to specify the same security group as the container group.
* `vswitch_id` - (Required, ForceNew) The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
* `acr_registry_info` - (Optional, ForceNew) The information about the Container Registry (ACR) Enterprise Edition instances used to pull images without a password. See [`acr_registry_info`](#acr_registry_info) below.
* `auto_match_image_cache` - (Optional, ForceNew) Specifies whether to enable the image cache auto match. Default to `false`.
* `eip_instance_id` - (Optional, ForceNew) The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
* `elimination_strategy` - (Optional, ForceNew) The elimination strategy of the image cache. Default to empty, which means the image cache is retained. Set the value to `LRU` to allow the image cache to be automatically deleted when the image cache quota is reached.
* `flash` - (Optional, ForceNew) Specifies whether to enable the fast image cache feature. The fast image cache feature accelerates the creation of image caches. Default to `false`. Note: Using this feature generates a temporary local snapshot and the snapshot fast availability fee is charged.
* `flash_copy_count` - (Optional, ForceNew) The number of local snapshot copies. We recommend that you add one snapshot copy for every 1000 ECI instances. When the fast image cache feature (`flash` is `true`) is enabled, a local snapshot is created first and deleted automatically after the normal snapshot is created.
* `image_cache_size`   - (Optional, ForceNew) The size of the image cache. Default to `20`. Unit: GiB.
* `insecure_registry` - (Optional, ForceNew) The address of the self-managed image repository. When you use an image in a self-managed repository that uses a self-signed certificate, you must configure this parameter to skip certificate authentication.
* `os_type` - (Optional, ForceNew) The operating system of the container image. Valid values: `Linux`, `Windows`. Default to `Linux`. Note: `Windows` is in invitational preview. To use it, submit a ticket.
* `plain_http_registry` - (Optional, ForceNew) The address of the self-managed image repository. When you use an image in a self-managed repository that uses the HTTP protocol, you must configure this parameter so that ECI pulls the image over HTTP.
* `resource_group_id` - (Optional, ForceNew) The ID of the resource group.
* `retention_days` - (Optional, ForceNew) The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
* `standard_copy_count` - (Optional, ForceNew) The number of normal snapshot copies. We recommend that you add one snapshot copy for every 1000 ECI instances. When the fast image cache feature (`flash` is `false`) is not enabled, only normal snapshots are generated.
* `zone_id` - (Optional, ForceNew) The zone id to cache image.
* `image_registry_credential` - (Optional, ForceNew) The Image Registry parameters about the image to be cached. See [`image_registry_credential`](#image_registry_credential) below.

### `acr_registry_info`
* `arn_service` - (Optional) The ARN of the RAM role in the account to which the ECI resources belong. It is used to specify the RAM role for pulling ACR images without a password.
* `arn_user` - (Optional) The ARN of the RAM role in the account to which the ACR instance belongs. It is used to pull ACR images across accounts without a password.
* `domain` - (Optional) The domain names of the ACR Enterprise Edition instance. By default, all domain names of the instance are used. Separate multiple domain names with commas (,).
* `instance_id` - (Optional) The ID of the ACR Enterprise Edition instance. This parameter is required when the ACR Enterprise Edition instance uses a custom domain name.
* `instance_name` - (Optional) The name of the ACR Enterprise Edition instance.
* `region_id` - (Optional) The region ID of the ACR Enterprise Edition instance. You need to specify this parameter when the ACR Enterprise Edition instance and the ECI resources are in different regions.

### `image_registry_credential`
* `server` - (Optional) The address of Image Registry without `http://` or `https://`.
* `user_name` - (Optional) The user name of Image Registry.
* `password` - (Optional) The password of the Image Registry.

## Attributes Reference

* `id` -The id of the image cache.
* `container_group_id` - The ID of the container group job that is used to create the image cache.
* `status` -The status of the image cache.

## Import

ECI Image Cache can be imported using the id, e.g.

```shell
$ terraform import alicloud_eci_image_cache.example abc123456
```
