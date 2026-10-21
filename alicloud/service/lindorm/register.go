package lindorm

import "github.com/aliyun/terraform-provider-alicloud/alicloud/provider/conns"

func ServicePackage() conns.ServicePackage {
	return conns.ServicePackage{
		Name: "Lindorm",
		SDKResources: []conns.SDKResource{
			{TypeName: "alicloud_lindorm_instance", Factory: ResourceInstance},
		},
	}
}
