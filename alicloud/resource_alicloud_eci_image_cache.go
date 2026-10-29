package alicloud

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/eci"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAlicloudEciImageCache() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudEciImageCacheCreate,
		Read:   resourceAlicloudEciImageCacheRead,
		Delete: resourceAlicloudEciImageCacheDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"acr_registry_info": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn_service": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"arn_user": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"domain": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"instance_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"region_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"auto_match_image_cache": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"container_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"eip_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"elimination_strategy": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"flash": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"flash_copy_count": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"image_cache_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"image_cache_size": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"image_registry_credential": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"server": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				ForceNew: true,
			},
			"images": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ForceNew: true,
			},
			"insecure_registry": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"os_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"Linux", "Windows"}, false),
			},
			"plain_http_registry": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"retention_days": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"security_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"standard_copy_count": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vswitch_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAlicloudEciImageCacheCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	eciService := EciService{client}

	request, err := client.NewCommonRequest("Eci", "Eci", strings.ToUpper(string(Https)), connectivity.ApiVersion("2018-08-08"))
	if err != nil {
		return WrapError(err)
	}
	request.ApiName = "CreateImageCache"
	request.QueryParams["RegionId"] = client.RegionId
	request.QueryParams["ImageCacheName"] = d.Get("image_cache_name").(string)
	request.QueryParams["SecurityGroupId"] = d.Get("security_group_id").(string)
	request.QueryParams["VSwitchId"] = d.Get("vswitch_id").(string)
	if v, ok := d.GetOk("eip_instance_id"); ok {
		request.QueryParams["EipInstanceId"] = v.(string)
	}
	if v, ok := d.GetOk("image_cache_size"); ok {
		request.QueryParams["ImageCacheSize"] = strconv.Itoa(v.(int))
	}
	if v, ok := d.GetOk("image_registry_credential"); ok {
		for i, e := range v.(*schema.Set).List() {
			credential := e.(map[string]interface{})
			index := strconv.Itoa(i + 1)
			if password, ok := credential["password"]; ok && password.(string) != "" {
				request.QueryParams["ImageRegistryCredential."+index+".Password"] = password.(string)
			}
			if server, ok := credential["server"]; ok && server.(string) != "" {
				request.QueryParams["ImageRegistryCredential."+index+".Server"] = server.(string)
			}
			if userName, ok := credential["user_name"]; ok && userName.(string) != "" {
				request.QueryParams["ImageRegistryCredential."+index+".UserName"] = userName.(string)
			}
		}
	}
	for i, image := range expandStringList(d.Get("images").(*schema.Set).List()) {
		request.QueryParams["Image."+strconv.Itoa(i+1)] = image
	}
	if v, ok := d.GetOk("acr_registry_info"); ok {
		for i, e := range v.(*schema.Set).List() {
			registry := e.(map[string]interface{})
			prefix := "AcrRegistryInfo." + strconv.Itoa(i+1)
			for key, param := range map[string]string{
				"arn_service":   "ArnService",
				"arn_user":      "ArnUser",
				"instance_id":   "InstanceId",
				"instance_name": "InstanceName",
				"region_id":     "RegionId",
			} {
				if value, ok := registry[key]; ok && value.(string) != "" {
					request.QueryParams[prefix+"."+param] = value.(string)
				}
			}
			if v, ok := registry["domain"]; ok {
				for j, domain := range v.(*schema.Set).List() {
					request.QueryParams[prefix+".Domain."+strconv.Itoa(j+1)] = domain.(string)
				}
			}
		}
	}
	if v, ok := d.GetOkExists("auto_match_image_cache"); ok {
		request.QueryParams["AutoMatchImageCache"] = strconv.FormatBool(v.(bool))
	}
	if v, ok := d.GetOkExists("flash"); ok {
		request.QueryParams["Flash"] = strconv.FormatBool(v.(bool))
	}
	if v, ok := d.GetOk("flash_copy_count"); ok {
		request.QueryParams["FlashCopyCount"] = strconv.Itoa(v.(int))
	}
	if v, ok := d.GetOk("standard_copy_count"); ok {
		request.QueryParams["StandardCopyCount"] = strconv.Itoa(v.(int))
	}
	if v, ok := d.GetOk("elimination_strategy"); ok {
		request.QueryParams["EliminationStrategy"] = v.(string)
	}
	if v, ok := d.GetOk("insecure_registry"); ok {
		request.QueryParams["InsecureRegistry"] = v.(string)
	}
	if v, ok := d.GetOk("plain_http_registry"); ok {
		request.QueryParams["PlainHttpRegistry"] = v.(string)
	}
	if v, ok := d.GetOk("os_type"); ok {
		request.QueryParams["OsType"] = v.(string)
	}
	if v, ok := d.GetOk("resource_group_id"); ok {
		request.QueryParams["ResourceGroupId"] = v.(string)
	}
	if v, ok := d.GetOk("retention_days"); ok {
		request.QueryParams["RetentionDays"] = strconv.Itoa(v.(int))
	}
	if v, ok := d.GetOk("zone_id"); ok {
		request.QueryParams["ZoneId"] = v.(string)
	}

	raw, err := client.WithEciClient(func(eciClient *eci.Client) (interface{}, error) {
		return eciClient.ProcessCommonRequest(request)
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_eci_image_cache", request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.Headers, request)
	response, _ := raw.(*responses.CommonResponse)
	var object map[string]interface{}
	if err := json.Unmarshal(response.GetHttpContentBytes(), &object); err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_eci_image_cache", request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	if _, ok := object["ImageCacheId"]; !ok {
		return WrapErrorf(fmt.Errorf("missing ImageCacheId in CreateImageCache response"), DefaultErrorMsg, "alicloud_eci_image_cache", request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	d.SetId(fmt.Sprint(object["ImageCacheId"]))
	stateConf := BuildStateConf([]string{"Preparing", "Creating"}, []string{"Ready"}, d.Timeout(schema.TimeoutCreate), 10*time.Second, eciService.EciImageCacheStateRefreshFunc(d.Id(), []string{"Failed"}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAlicloudEciImageCacheRead(d, meta)
}
func resourceAlicloudEciImageCacheRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	eciService := EciService{client}
	object, err := eciService.DescribeEciImageCache(d.Id())
	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("container_group_id", object.ContainerGroupId)
	d.Set("image_cache_name", object.ImageCacheName)
	d.Set("images", object.Images)
	d.Set("status", object.Status)
	return nil
}
func resourceAlicloudEciImageCacheDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	request := eci.CreateDeleteImageCacheRequest()
	request.ImageCacheId = d.Id()
	request.RegionId = client.RegionId
	raw, err := client.WithEciClient(func(eciClient *eci.Client) (interface{}, error) {
		return eciClient.DeleteImageCache(request)
	})
	addDebug(request.GetActionName(), raw)
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	return nil
}
