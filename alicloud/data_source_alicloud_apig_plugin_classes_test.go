package alicloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceAliCloudApigPluginClasses_basic(t *testing.T) {
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccapig%d", rand)
	resourceId := "data.alicloud_apig_plugin_classes.default"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAliCloudApigPluginClassesConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceId, "plugin_classes.#"),
				),
			},
		},
	})
}

func testAccDataSourceAliCloudApigPluginClassesConfig(name string) string {
	return fmt.Sprintf(`
resource "alicloud_apig_plugin_class" "default" {
  wasm_url                      = "https://example.com/plugin.wasm"
  description                   = "A test plugin class for data source"
  version_description           = "Initial version"
  plugin_class_name             = "%[1]s"
  version                       = "1.0.0"
  execute_priority              = 1
  wasm_language                 = "TinyGo"
  execute_stage                 = "UNSPECIFIED_PHASE"
}

data "alicloud_apig_plugin_classes" "default" {
  ids = [alicloud_apig_plugin_class.default.id]
}
`, name)
}
