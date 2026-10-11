// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudPaiLangStudioToolset() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPaiLangStudioToolsetCreate,
		Read:   resourceAliCloudPaiLangStudioToolsetRead,
		Update: resourceAliCloudPaiLangStudioToolsetUpdate,
		Delete: resourceAliCloudPaiLangStudioToolsetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"accessibility": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"creator": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
			},
			"toolset_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"toolset_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"toolset_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"workspace_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudPaiLangStudioToolsetCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/api/v1/langstudio/toolsets")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("workspace_id"); ok {
		request["WorkspaceId"] = v
	}

	if v, ok := d.GetOk("accessibility"); ok {
		request["Accessibility"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("toolset_config"); ok {
		request["ToolsetConfig"] = v
	}
	if v, ok := d.GetOk("toolset_name"); ok {
		request["ToolsetName"] = v
	}
	if v, ok := d.GetOk("toolset_type"); ok {
		request["ToolsetType"] = v
	}
	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RoaPost("PAILangStudio", "2024-07-10", action, query, nil, body, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_pai_lang_studio_toolset", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["WorkspaceId"], response["ToolsetId"]))

	return resourceAliCloudPaiLangStudioToolsetUpdate(d, meta)
}

func resourceAliCloudPaiLangStudioToolsetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	paiLangStudioServiceV2 := PaiLangStudioServiceV2{client}

	objectRaw, err := paiLangStudioServiceV2.DescribePaiLangStudioToolset(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_pai_lang_studio_toolset DescribePaiLangStudioToolset Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("accessibility", objectRaw["Accessibility"])
	d.Set("creator", objectRaw["Creator"])
	d.Set("description", objectRaw["Description"])
	d.Set("gmt_create_time", objectRaw["GmtCreateTime"])
	d.Set("gmt_modified_time", objectRaw["GmtModifiedTime"])
	d.Set("tools", objectRaw["Tools"])
	d.Set("toolset_config", objectRaw["ToolsetConfig"])
	d.Set("toolset_name", objectRaw["ToolsetName"])
	d.Set("toolset_type", objectRaw["ToolsetType"])
	d.Set("toolset_id", objectRaw["ToolsetId"])
	d.Set("workspace_id", objectRaw["WorkspaceId"])

	return nil
}

func resourceAliCloudPaiLangStudioToolsetUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	ToolsetId := parts[1]
	action := fmt.Sprintf("/api/v1/langstudio/toolsets/%s", ToolsetId)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	request["WorkspaceId"] = parts[0]

	if !d.IsNewResource() && d.HasChange("description") {
		update = true
	}
	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		request["Description"] = v
	}
	if !d.IsNewResource() && d.HasChange("toolset_config") {
		update = true
	}
	if v, ok := d.GetOk("toolset_config"); ok || d.HasChange("toolset_config") {
		request["ToolsetConfig"] = v
	}
	if !d.IsNewResource() && d.HasChange("toolset_name") {
		update = true
	}
	if v, ok := d.GetOk("toolset_name"); ok || d.HasChange("toolset_name") {
		request["ToolsetName"] = v
	}
	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaPut("PAILangStudio", "2024-07-10", action, query, nil, body, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}

	return resourceAliCloudPaiLangStudioToolsetRead(d, meta)
}

func resourceAliCloudPaiLangStudioToolsetDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	ToolsetId := parts[1]
	action := fmt.Sprintf("/api/v1/langstudio/toolsets/%s", ToolsetId)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	query["WorkspaceId"] = StringPointer(parts[0])

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RoaDelete("PAILangStudio", "2024-07-10", action, query, nil, nil, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
