package main

import (
  "github.com/FireDrunk/terraform-proxmox/proxmox"
  "github.com/hashicorp/terraform/plugin"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{ProviderFunc: proxmoxprovider.Provider})
}
