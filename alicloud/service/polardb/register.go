package polardb

import (
	"github.com/aliyun/terraform-provider-alicloud/alicloud/provider/conns"
)

// ServicePackage returns the registration declaration of the PolarDB service.
func ServicePackage() conns.ServicePackage {
	return conns.ServicePackage{
		Name: "Polardb",
		SDKResources: []conns.SDKResource{
			{
				TypeName: "alicloud_polardb_aidb_cluster_dataset",
				Factory:  ResourceAIDBClusterDataset,
			},
		},
	}
}
