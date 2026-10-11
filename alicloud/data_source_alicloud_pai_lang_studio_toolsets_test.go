package alicloud

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

func TestAccAliCloudPaiLangStudioToolsetsDataSource(t *testing.T) {
	testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
	rand := acctest.RandIntRange(10000, 99999)

	idsConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAliCloudPaiLangStudioToolsetsSourceConfig(rand, map[string]string{
			"workspace_id": `"${alicloud_pai_workspace_workspace.default.id}"`,
			"ids":          `["${alicloud_pai_lang_studio_toolset.default.id}"]`,
		}),
		fakeConfig: testAccCheckAliCloudPaiLangStudioToolsetsSourceConfig(rand, map[string]string{
			"workspace_id": `"${alicloud_pai_workspace_workspace.default.id}"`,
			"ids":          `["${alicloud_pai_lang_studio_toolset.default.id}_fake"]`,
		}),
	}

	nameRegexConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAliCloudPaiLangStudioToolsetsSourceConfig(rand, map[string]string{
			"workspace_id": `"${alicloud_pai_workspace_workspace.default.id}"`,
			"name_regex":   `"${alicloud_pai_lang_studio_toolset.default.toolset_name}"`,
		}),
		fakeConfig: testAccCheckAliCloudPaiLangStudioToolsetsSourceConfig(rand, map[string]string{
			"workspace_id": `"${alicloud_pai_workspace_workspace.default.id}"`,
			"name_regex":   `"${alicloud_pai_lang_studio_toolset.default.toolset_name}_fake"`,
		}),
	}

	allFiltersConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAliCloudPaiLangStudioToolsetsSourceConfig(rand, map[string]string{
			"workspace_id":  `"${alicloud_pai_workspace_workspace.default.id}"`,
			"accessibility": `"PRIVATE"`,
			"creator":       `"${alicloud_pai_lang_studio_toolset.default.creator}"`,
			"toolset_id":    `"${alicloud_pai_lang_studio_toolset.default.toolset_id}"`,
			"toolset_name":  `"${alicloud_pai_lang_studio_toolset.default.toolset_name}"`,
			"toolset_type":  `"OPENAPI"`,
			"order":         `"desc"`,
			"sort_by":       `"gmtCreateTime"`,
		}),
		fakeConfig: testAccCheckAliCloudPaiLangStudioToolsetsSourceConfig(rand, map[string]string{
			"workspace_id":  `"${alicloud_pai_workspace_workspace.default.id}"`,
			"accessibility": `"PRIVATE"`,
			"toolset_name":  `"${alicloud_pai_lang_studio_toolset.default.toolset_name}_fake"`,
		}),
	}

	PaiLangStudioToolsetsCheckInfo.dataSourceTestCheck(t, rand, idsConf, nameRegexConf, allFiltersConf)
}

var existPaiLangStudioToolsetsMapFunc = func(rand int) map[string]string {
	return map[string]string{
		"ids.#":                        "1",
		"names.#":                      "1",
		"toolsets.#":                   "1",
		"toolsets.0.id":                CHECKSET,
		"toolsets.0.workspace_id":      CHECKSET,
		"toolsets.0.toolset_id":        CHECKSET,
		"toolsets.0.toolset_name":      fmt.Sprintf("tfaccpailangstudio%d", rand),
		"toolsets.0.toolset_type":      "OPENAPI",
		"toolsets.0.accessibility":     "PRIVATE",
		"toolsets.0.description":       fmt.Sprintf("tfaccpailangstudio%d", rand),
		"toolsets.0.creator":           CHECKSET,
		"toolsets.0.gmt_create_time":   CHECKSET,
		"toolsets.0.gmt_modified_time": CHECKSET,
	}
}

var fakePaiLangStudioToolsetsMapFunc = func(rand int) map[string]string {
	return map[string]string{
		"ids.#":      "0",
		"toolsets.#": "0",
	}
}

var PaiLangStudioToolsetsCheckInfo = dataSourceAttr{
	resourceId:   "data.alicloud_pai_lang_studio_toolsets.default",
	existMapFunc: existPaiLangStudioToolsetsMapFunc,
	fakeMapFunc:  fakePaiLangStudioToolsetsMapFunc,
}

func testAccCheckAliCloudPaiLangStudioToolsetsSourceConfig(rand int, attrMap map[string]string) string {
	var pairs []string
	for k, v := range attrMap {
		pairs = append(pairs, k+" = "+v)
	}
	name := fmt.Sprintf("tfaccpailangstudio%d", rand)
	return AliCloudPaiLangStudioToolsetBasicDependence0(name) + fmt.Sprintf(`
resource "alicloud_pai_lang_studio_toolset" "default" {
  workspace_id   = alicloud_pai_workspace_workspace.default.id
  toolset_name   = "%s"
  toolset_type   = "OPENAPI"
  accessibility  = "PRIVATE"
  description    = "%s"
  toolset_config = jsonencode({openapi="3.0.0",info={title="acc",version="1.0.0"},paths={"/acc"={get={operationId="acc",responses={"200"={description="ok"}}}}}})
}

data "alicloud_pai_lang_studio_toolsets" "default" {
  %s
}
`, name, name, strings.Join(pairs, "\n  "))
}
