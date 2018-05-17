package proxmoxprovider

import (
  "fmt"

	"github.com/hashicorp/terraform/helper/schema"
  "github.com/FireDrunk/go-proxmox"
)

func dataSourceProxmoxResourcePool() *schema.Resource {
	return &schema.Resource{
		Create:     dataSourceProxmoxResourcePoolCreate,
    Read:       dataSourceProxmoxResourcePoolRead,
    Update:     dataSourceProxmoxResourcePoolCreate,
    Delete:     dataSourceProxmoxResourcePoolDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the resource pool.",
        Required:    true,
				Optional:    false,
			},
      "comment": {
				Type:        schema.TypeString,
				Description: "The comment of the resource pool.",
				Optional:    true,
			},
		},
	}
}

func dataSourceProxmoxResourcePoolCreate(d *schema.ResourceData, _ interface{}) error {
  name := d.Get("name").(string)
  comment := d.Get("comment").(string)

  proxmoxClient, err := proxmox.NewProxMox("10.0.2.15:8006", "root", "password")
  if err != nil {
    fmt.Printf("Error: %s", err)
    return err
  }

  result, err := proxmoxClient.NewPool(name, comment)
  if err != nil {
    fmt.Printf("Error: %s", err)
    return err
  }

  fmt.Printf("Result: %s", result)

  return nil
}

func dataSourceProxmoxResourcePoolRead(d *schema.ResourceData, meta interface{}) error {
  name := d.Get("name").(string)

  proxmoxClient, error := proxmox.NewProxMox("10.0.2.15:8006", "root", "password")
  if error != nil {
    return error
  }

  pools, error := proxmoxClient.Pools()
  if error != nil {
    return error
  }

  for _, pool := range pools {
    if pool.Poolid == name {
      d.SetId(name)
    }
  }
  return nil
}

func dataSourceProxmoxResourcePoolDelete(d *schema.ResourceData, meta interface{}) error {
  name := d.Get("name").(string)

  proxmoxClient, err := proxmox.NewProxMox("10.0.2.15:8006", "root", "password")
  if err != nil {
    return err
  }

  poolerr := proxmoxClient.DeletePool(name)
  if poolerr != nil {
    return poolerr
  }

  return nil
}
