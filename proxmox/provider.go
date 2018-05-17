package proxmoxprovider

import (
  //builtin go packages
  "fmt"

  //public packages
  "github.com/hashicorp/terraform/helper/schema"
  "github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
  fmt.Println("initializing provider")
  return &schema.Provider{
		Schema: map[string]*schema.Schema{},

		ResourcesMap: map[string]*schema.Resource{
		  "proxmox_resource_pool": dataSourceProxmoxResourcePool(),
		},

		// DataSourcesMap: map[string]*schema.Resource{
		//
		// },
	}
}
