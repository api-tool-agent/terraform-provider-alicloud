package cloudfirewall

import (
	"github.com/aliyun/terraform-provider-alicloud/alicloud/provider/conns"
)

// ServicePackage returns the registration declaration of the Cloud Firewall service.
func ServicePackage() conns.ServicePackage {
	return conns.ServicePackage{
		Name: "CloudFirewall",
		SDKResources: []conns.SDKResource{
			{
				TypeName: "alicloud_cloud_firewall_address_book",
				Factory:  ResourceAddressBook,
			},
		},
	}
}
