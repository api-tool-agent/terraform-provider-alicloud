package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAliCloudPaiLangStudioToolset_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_pai_lang_studio_toolset.default"
	ra := resourceAttrInit(resourceId, AliCloudPaiLangStudioToolsetMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PaiLangStudioServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePaiLangStudioToolset")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccpailangstudio%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudPaiLangStudioToolsetBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"workspace_id":   "${alicloud_pai_workspace_workspace.default.id}",
					"toolset_name":   name,
					"toolset_type":   "OPENAPI",
					"accessibility":  "PRIVATE",
					"description":    name,
					"toolset_config": "{\"openapi\":\"3.0.0\",\"info\":{\"title\":\"acc\",\"version\":\"1.0.0\"},\"paths\":{\"/acc\":{\"get\":{\"operationId\":\"acc\",\"responses\":{\"200\":{\"description\":\"ok\"}}}}}}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"workspace_id":   CHECKSET,
						"toolset_name":   name,
						"toolset_type":   "OPENAPI",
						"accessibility":  "PRIVATE",
						"description":    name,
						"toolset_config": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "_update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"toolset_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"toolset_name": name + "_update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"toolset_config": "{\"openapi\":\"3.0.0\",\"info\":{\"title\":\"acc2\",\"version\":\"1.0.0\"},\"paths\":{\"/acc\":{\"get\":{\"operationId\":\"acc\",\"responses\":{\"200\":{\"description\":\"ok\"}}}}}}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"toolset_config": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAliCloudPaiLangStudioToolset_basic0_twin(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_pai_lang_studio_toolset.default"
	ra := resourceAttrInit(resourceId, AliCloudPaiLangStudioToolsetMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PaiLangStudioServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePaiLangStudioToolset")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccpailangstudio%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudPaiLangStudioToolsetBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"workspace_id":   "${alicloud_pai_workspace_workspace.default.id}",
					"toolset_name":   name,
					"toolset_type":   "OPENAPI",
					"accessibility":  "PUBLIC",
					"description":    name,
					"toolset_config": "{\"openapi\":\"3.0.0\",\"info\":{\"title\":\"acc\",\"version\":\"1.0.0\"},\"paths\":{\"/acc\":{\"get\":{\"operationId\":\"acc\",\"responses\":{\"200\":{\"description\":\"ok\"}}}}}}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"workspace_id":   CHECKSET,
						"toolset_name":   name,
						"toolset_type":   "OPENAPI",
						"accessibility":  "PUBLIC",
						"description":    name,
						"toolset_config": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AliCloudPaiLangStudioToolsetMap0 = map[string]string{}

func AliCloudPaiLangStudioToolsetBasicDependence0(name string) string {
	return fmt.Sprintf(`
variable "name" {
  default = "%s"
}

resource "alicloud_pai_workspace_workspace" "default" {
  description    = "for pai lang studio toolset acc test"
  display_name   = var.name
  workspace_name = var.name
  env_types      = ["prod"]
}
`, name)
}
