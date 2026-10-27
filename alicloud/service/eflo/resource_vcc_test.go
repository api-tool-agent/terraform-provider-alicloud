// Package eflo. This file is generated automatically. Please do not modify it manually, thank you!
package eflo_test

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/acctest"
	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// Test Eflo Vcc. >>> Resource test cases, automatically generated.
// Case vcc-test-vpc_1 12634
func TestAccAliCloudEfloVcc_basic12634(t *testing.T) {
	r := tfacctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceflo%d", r)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			t.Setenv("ALICLOUD_REGION", "cn-wulanchabu")
			acctest.PreCheck(t)
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAliCloudEfloVcc_basic12634_step0(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "description", "test-tf"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "access_could_service", "false"),

					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "resource_group_id"),
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "connection_type", "VPC"),

					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "zone_id"),
					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "vswitch_id"),
					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "vpc_id"),
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "bgp_asn", "64512"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "bgp_cidr", "169.254.128.0/23"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "bandwidth", "1000"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "vcc_name", name),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("alicloud_eflo_vcc.default", tfjsonpath.New("create_time"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_eflo_vcc.default", tfjsonpath.New("gmt_modified"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_eflo_vcc.default", tfjsonpath.New("message"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_eflo_vcc.default", tfjsonpath.New("spec"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("alicloud_eflo_vcc.default", tfjsonpath.New("status"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccAliCloudEfloVcc_basic12634_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "description", "test-tf"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "access_could_service", "false"),

					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "resource_group_id"),
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "connection_type", "VPC"),

					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "zone_id"),
					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "vswitch_id"),
					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "vpc_id"),
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "bgp_asn", "64512"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "bgp_cidr", "169.254.128.0/23"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "bandwidth", "2000"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "vcc_name", name+"-update"),
				),
			},
			{
				Config: testAccAliCloudEfloVcc_basic12634_step2(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "description", "test-tf"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "access_could_service", "false"),

					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "resource_group_id"),
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "connection_type", "CENTR"),

					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "zone_id"),
					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "vswitch_id"),
					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "vpc_id"),
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "bgp_asn", "64512"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "bgp_cidr", "169.254.128.0/23"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "bandwidth", "2000"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "vcc_name", name+"-update"),
					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "cen_id"),
				),
			},
			{
				Config: testAccAliCloudEfloVcc_basic12634_step3(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "description", "test-tf"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "access_could_service", "false"),

					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "resource_group_id"),
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "connection_type", "VPC"),

					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "zone_id"),
					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "vswitch_id"),
					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "vpc_id"),
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "bgp_asn", "64512"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "bgp_cidr", "169.254.128.0/23"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "bandwidth", "2000"),

					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "vcc_name", name+"-update"),
					resource.TestCheckResourceAttrSet("alicloud_eflo_vcc.default", "cen_id"),
				),
			},
			{
				Config: testAccAliCloudEfloVcc_basic12634_tags_step1(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "tags.%", "2"),
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "tags.Created", "TF"),
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "tags.For", "Test"),
				),
			},
			{
				Config: testAccAliCloudEfloVcc_basic12634_tags_step2(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "tags.%", "2"),
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "tags.Created", "TF-update"),
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "tags.For", "Test-update"),
				),
			},
			{
				Config: testAccAliCloudEfloVcc_basic12634_tags_step3(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_eflo_vcc.default", "tags.%", "0"),
				),
			},
			{
				ResourceName:            "alicloud_eflo_vcc.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"access_could_service", "bgp_asn", "description"},
			},
		},
	})
}

func testAccAliCloudEfloVcc_basic12634_base(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%[1]s"
}

variable "region_id" {
  default = "cn-wulanchabu"
}

variable "zone_id" {
  default = "cn-wulanchabu-b"
}

resource "alicloud_vpc" "VPC" {
  is_default  = false
  description = "vcc测试用"
  cidr_block  = "192.168.0.0/16"
  enable_ipv6 = false
  vpc_name    = "vpc-vcc-test-tf"
}

resource "alicloud_vswitch" "VSW" {
  zone_id      = var.zone_id
  cidr_block   = "192.168.0.0/24"
  vswitch_name = "vsw-vcc-tf"
  vpc_id       = alicloud_vpc.VPC.id
}

resource "alicloud_cen_instance" "CEN" {
  cen_instance_name = "cen-tf-test-1"
}

resource "alicloud_cen_transit_router" "CEN-TR" {
  cen_id              = alicloud_cen_instance.CEN.id
  transit_router_name = "tf-test"
}

resource "alicloud_cen_transit_router_vpc_attachment" "VPC-ATTACH" {
  vpc_id                             = alicloud_vpc.VPC.id
  transit_router_vpc_attachment_name = "tf-test"
  cen_id                             = alicloud_cen_instance.CEN.id
  zone_mappings {
    vswitch_id = alicloud_vswitch.VSW.id
    zone_id    = var.zone_id
  }
  transit_router_id = alicloud_cen_transit_router.CEN-TR.transit_router_id
}

resource "alicloud_vpc" "VPC-CRAETE" {
  is_default  = false
  description = "TF测试"
  cidr_block  = "192.168.0.0/16"
  enable_ipv6 = false
  vpc_name    = "vpc-vcc-test-tf-2"
}

resource "alicloud_vswitch" "VSW-CREATE" {
  vpc_id       = alicloud_vpc.VPC-CRAETE.id
  zone_id      = var.zone_id
  cidr_block   = "192.168.0.0/24"
  vswitch_name = "vsw-vcc-tf-2"
}

resource "alicloud_resource_manager_resource_group" "RG" {
  resource_group_name = "tf-acc-eflo-vcc-${var.name}"
  display_name        = "tf-acc-eflo-vcc-${var.name}"
}

resource "alicloud_resource_manager_resource_group" "RG2" {
  resource_group_name = "tf-acc-eflo-vcc2-${var.name}"
  display_name        = "tf-acc-eflo-vcc2-${var.name}"
}


`, name)
}

func testAccAliCloudEfloVcc_basic12634_step0(name string) string {
	return fmt.Sprintf(`
%[2]s

resource "alicloud_eflo_vcc" "default" {
        description = "test-tf"

        access_could_service = false

        resource_group_id = "${alicloud_resource_manager_resource_group.RG.id}"

        connection_type = "VPC"

        zone_id = "${var.zone_id}"

        vswitch_id = "${alicloud_vswitch.VSW-CREATE.id}"

        vpc_id = "${alicloud_vpc.VPC-CRAETE.id}"

        bgp_asn = 64512

        bgp_cidr = "169.254.128.0/23"

        bandwidth = 1000

        vcc_name = "%[1]s"
}
`, name, testAccAliCloudEfloVcc_basic12634_base(name))
}

func testAccAliCloudEfloVcc_basic12634_step1(name string) string {
	return fmt.Sprintf(`
%[2]s

resource "alicloud_eflo_vcc" "default" {
        description = "test-tf"

        access_could_service = false

        resource_group_id = "${alicloud_resource_manager_resource_group.RG2.id}"

        connection_type = "VPC"

        zone_id = "${var.zone_id}"

        vswitch_id = "${alicloud_vswitch.VSW-CREATE.id}"

        vpc_id = "${alicloud_vpc.VPC-CRAETE.id}"

        bgp_asn = 64512

        bgp_cidr = "169.254.128.0/23"

        bandwidth = 2000

        vcc_name = "%[1]s-update"
}
`, name, testAccAliCloudEfloVcc_basic12634_base(name))
}

func testAccAliCloudEfloVcc_basic12634_step2(name string) string {
	return fmt.Sprintf(`
%[2]s

resource "alicloud_eflo_vcc" "default" {
        description = "test-tf"

        access_could_service = false

        resource_group_id = "${alicloud_resource_manager_resource_group.RG2.id}"

        connection_type = "CENTR"

        zone_id = "${var.zone_id}"

        vswitch_id = "${alicloud_vswitch.VSW.id}"

        vpc_id = "${alicloud_vpc.VPC.id}"

        bgp_asn = 64512

        bgp_cidr = "169.254.128.0/23"

        bandwidth = 2000

        vcc_name = "%[1]s-update"
        cen_id = "${alicloud_cen_instance.CEN.id}"

}
`, name, testAccAliCloudEfloVcc_basic12634_base(name))
}

func testAccAliCloudEfloVcc_basic12634_step3(name string) string {
	return fmt.Sprintf(`
%[2]s

resource "alicloud_eflo_vcc" "default" {
        description = "test-tf"

        access_could_service = false

        resource_group_id = "${alicloud_resource_manager_resource_group.RG2.id}"

        connection_type = "VPC"

        zone_id = "${var.zone_id}"

        vswitch_id = "${alicloud_vswitch.VSW-CREATE.id}"

        vpc_id = "${alicloud_vpc.VPC-CRAETE.id}"

        bgp_asn = 64512

        bgp_cidr = "169.254.128.0/23"

        bandwidth = 2000

        vcc_name = "%[1]s-update"
        cen_id = "${alicloud_cen_instance.CEN.id}"

}
`, name, testAccAliCloudEfloVcc_basic12634_base(name))
}

func testAccAliCloudEfloVcc_basic12634_tags_step1(name string) string {
	return fmt.Sprintf(`
%[2]s

resource "alicloud_eflo_vcc" "default" {
        description = "test-tf"

        access_could_service = false

        resource_group_id = "${alicloud_resource_manager_resource_group.RG2.id}"

        connection_type = "VPC"

        zone_id = "${var.zone_id}"

        vswitch_id = "${alicloud_vswitch.VSW-CREATE.id}"

        vpc_id = "${alicloud_vpc.VPC-CRAETE.id}"

        bgp_asn = 64512

        bgp_cidr = "169.254.128.0/23"

        bandwidth = 2000

        vcc_name = "%[1]s-update"
        cen_id = "${alicloud_cen_instance.CEN.id}"

    tags = {
        Created = "TF"
        For     = "Test"
    }
}
`, name, testAccAliCloudEfloVcc_basic12634_base(name))
}

func testAccAliCloudEfloVcc_basic12634_tags_step2(name string) string {
	return fmt.Sprintf(`
%[2]s

resource "alicloud_eflo_vcc" "default" {
        description = "test-tf"

        access_could_service = false

        resource_group_id = "${alicloud_resource_manager_resource_group.RG2.id}"

        connection_type = "VPC"

        zone_id = "${var.zone_id}"

        vswitch_id = "${alicloud_vswitch.VSW-CREATE.id}"

        vpc_id = "${alicloud_vpc.VPC-CRAETE.id}"

        bgp_asn = 64512

        bgp_cidr = "169.254.128.0/23"

        bandwidth = 2000

        vcc_name = "%[1]s-update"
        cen_id = "${alicloud_cen_instance.CEN.id}"

    tags = {
        Created = "TF-update"
        For     = "Test-update"
    }
}
`, name, testAccAliCloudEfloVcc_basic12634_base(name))
}

func testAccAliCloudEfloVcc_basic12634_tags_step3(name string) string {
	return fmt.Sprintf(`
%[2]s

resource "alicloud_eflo_vcc" "default" {
        description = "test-tf"

        access_could_service = false

        resource_group_id = "${alicloud_resource_manager_resource_group.RG2.id}"

        connection_type = "VPC"

        zone_id = "${var.zone_id}"

        vswitch_id = "${alicloud_vswitch.VSW-CREATE.id}"

        vpc_id = "${alicloud_vpc.VPC-CRAETE.id}"

        bgp_asn = 64512

        bgp_cidr = "169.254.128.0/23"

        bandwidth = 2000

        vcc_name = "%[1]s-update"
        cen_id = "${alicloud_cen_instance.CEN.id}"

  tags = {}
}
`, name, testAccAliCloudEfloVcc_basic12634_base(name))
}
