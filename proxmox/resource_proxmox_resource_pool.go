package proxmoxprovider

import (
  "fmt"
  "log"

	"github.com/hashicorp/terraform/helper/schema"
  "github.com/FireDrunk/go-proxmox"
)

func resourceProxmoxResourcePool() *schema.Resource {
	return &schema.Resource{
		Create:     resourceProxmoxResourcePoolCreate,
    Read:       resourceProxmoxResourcePoolRead,
    Update:     resourceProxmoxResourcePoolUpdate,
    Delete:     resourceProxmoxResourcePoolDelete,

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

func resourceProxmoxResourcePoolCreate(d *schema.ResourceData, _ interface{}) error {
  name := d.Get("name").(string)
  comment := d.Get("comment").(string)

  log.Print("Create method executed!")

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

func resourceProxmoxResourcePoolRead(d *schema.ResourceData, _ interface{}) error {
  log.Print("Read method executed!")

  var found bool

  name := d.Get("name").(string)

  proxmoxClient, error := proxmox.NewProxMox("10.0.2.15:8006", "root", "password")
  if error != nil {
    return error
  }

  log.Fatal("proxmox-resource-pool-read called")

  pools, err := proxmoxClient.Pools()
  if err != nil {
    return fmt.Errorf("Error retrieving pools: %s", err)
  }

  for _, pool := range pools {
    if pool.Poolid == name {
      found = true
      d.SetId(pool.Poolid)
    }
  }

  if !found {
    d.SetId("")
  }

  return nil
}

func resourceProxmoxResourcePoolUpdate(d *schema.ResourceData, _ interface{}) error {
  name := d.Get("name").(string)
  comment := d.Get("comment").(string)

  proxmoxClient, err := proxmox.NewProxMox("10.0.2.15:8006", "root", "password")
  if err != nil {
    fmt.Printf("Error: %s", err)
    return err
  }

  result, err := proxmoxClient.UpdatePool(name, comment)
  if err != nil {
    fmt.Printf("Error: %s", err)
    return err
  }

  fmt.Printf("Result: %s", result)

  return nil
}

func resourceProxmoxResourcePoolDelete(d *schema.ResourceData, meta interface{}) error {
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
