package cms

import (
	"github.com/aliyun/terraform-provider-alicloud/alicloud/provider/conns"
)

// ServicePackage returns the registration declaration of the CMS service.
func ServicePackage() conns.ServicePackage {
	return conns.ServicePackage{
		Name: "Cms",
		SDKResources: []conns.SDKResource{
			{
				TypeName: "alicloud_cms_alert_rule_v2",
				Factory:  ResourceAlertRuleV2,
			},
		},
		DataSources: []conns.DataSource{
			{
				TypeName: "alicloud_cms_metric_list",
				Factory:  NewMetricListDataSource,
			},
		},
	}
}
