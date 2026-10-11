// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAliCloudPaiLangStudioToolsets() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAliCloudPaiLangStudioToolsetRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"name_regex": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"names": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"accessibility": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"creator": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"order": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"toolset_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"toolset_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"toolset_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"workspace_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"toolsets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accessibility": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"creator": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gmt_create_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gmt_modified_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tools": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"toolset_config": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"toolset_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"toolset_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"toolset_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"workspace_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"output_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceAliCloudPaiLangStudioToolsetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	var objects []map[string]interface{}
	var nameRegex *regexp.Regexp
	if v, ok := d.GetOk("name_regex"); ok {
		r, err := regexp.Compile(v.(string))
		if err != nil {
			return WrapError(err)
		}
		nameRegex = r
	}

	idsMap := make(map[string]string)
	if v, ok := d.GetOk("ids"); ok {
		for _, vv := range v.([]interface{}) {
			if vv == nil {
				continue
			}
			idsMap[vv.(string)] = vv.(string)
		}
	}

	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	// ListToolsets
	action := fmt.Sprintf("/api/v1/langstudio/toolsets")
	var err error
	request = make(map[string]interface{})
	query = make(map[string]*string)

	if v, ok := d.GetOk("accessibility"); ok {
		query["Accessibility"] = StringPointer(v.(string))
	}

	if v, ok := d.GetOk("creator"); ok {
		query["Creator"] = StringPointer(v.(string))
	}

	if v, ok := d.GetOk("order"); ok {
		query["Order"] = StringPointer(v.(string))
	}

	if v, ok := d.GetOk("sort_by"); ok {
		query["SortBy"] = StringPointer(v.(string))
	}

	if v, ok := d.GetOk("toolset_id"); ok {
		query["ToolsetId"] = StringPointer(v.(string))
	}

	if v, ok := d.GetOk("toolset_name"); ok {
		query["ToolsetName"] = StringPointer(v.(string))
	}

	if v, ok := d.GetOk("toolset_type"); ok {
		query["ToolsetType"] = StringPointer(v.(string))
	}

	if v, ok := d.GetOk("workspace_id"); ok {
		query["WorkspaceId"] = StringPointer(v.(string))
	}

	query["MaxResults"] = StringPointer(strconv.Itoa(PageSizeLarge))
	for {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutRead), func() *resource.RetryError {
			response, err = client.RoaGet("PAILangStudio", "2024-07-10", action, query, nil, nil)

			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, response, request)
			return nil
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}

		resp, _ := jsonpath.Get("$.Toolsets[*]", response)

		result, _ := resp.([]interface{})
		for _, v := range result {
			item := v.(map[string]interface{})
			if nameRegex != nil && !nameRegex.MatchString(fmt.Sprint(item["ToolsetName"])) {
				continue
			}
			if len(idsMap) > 0 {
				if _, ok := idsMap[fmt.Sprint(item["WorkspaceId"], ":", item["ToolsetId"])]; !ok {
					continue
				}
			}
			objects = append(objects, item)
		}

		if nextToken, ok := response["NextToken"].(string); ok && nextToken != "" {
			query["NextToken"] = StringPointer(nextToken)
		} else {
			break
		}
	}

	ids := make([]string, 0)
	names := make([]interface{}, 0)
	s := make([]map[string]interface{}, 0)
	for _, objectRaw := range objects {
		mapping := map[string]interface{}{}

		mapping["id"] = fmt.Sprint(objectRaw["WorkspaceId"], ":", objectRaw["ToolsetId"])

		mapping["accessibility"] = objectRaw["Accessibility"]
		mapping["creator"] = objectRaw["Creator"]
		mapping["description"] = objectRaw["Description"]
		mapping["gmt_create_time"] = objectRaw["GmtCreateTime"]
		mapping["gmt_modified_time"] = objectRaw["GmtModifiedTime"]
		mapping["tools"] = objectRaw["Tools"]
		mapping["toolset_config"] = objectRaw["ToolsetConfig"]
		mapping["toolset_name"] = objectRaw["ToolsetName"]
		mapping["toolset_type"] = objectRaw["ToolsetType"]
		mapping["toolset_id"] = objectRaw["ToolsetId"]
		mapping["workspace_id"] = objectRaw["WorkspaceId"]

		ids = append(ids, fmt.Sprint(mapping["id"]))
		names = append(names, objectRaw["ToolsetName"])
		s = append(s, mapping)
	}

	d.SetId(dataResourceIdHash(ids))
	if err := d.Set("ids", ids); err != nil {
		return WrapError(err)
	}

	if err := d.Set("names", names); err != nil {
		return WrapError(err)
	}
	if err := d.Set("toolsets", s); err != nil {
		return WrapError(err)
	}

	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		writeToFile(output.(string), s)
	}
	return nil
}
