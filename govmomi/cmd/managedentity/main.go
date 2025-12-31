package main

import (
	"context"
	"fmt"
	"os"

	px "github.com/9506hqwy/samples-go/pkg/propertyex"
	sx "github.com/9506hqwy/samples-go/pkg/sessionex"
)

func main() {
	ctx := context.Background()

	url := os.Getenv("VSPHERE_URL")
	user := os.Getenv("VSPHERE_USER")
	pass := os.Getenv("VSPHERE_PASSWORD")

	c, err := sx.Login(ctx, url, user, pass)
	if err != nil {
		panic(err)
	}

	defer sx.Logout(ctx, c)

	moTypes := []string{
		"ClusterComputeResource",
		"ComputeResource",
		"Datacenter",
		"Datastore",
		"DistributedVirtualSwitch",
		"Folder",
		"HostSystem",
		"Network",
		"ResourcePool",
		"StoragePod",
		"VirtualApp",
		"VirtualMachine",
		"VmwareDistributedVirtualSwitch",
	}

	objects, err := px.RetrieveFromRoot(ctx, c, moTypes, []string{"name"})
	if err != nil {
		panic(err)
	}

	for _, obj := range objects {
		for _, prop := range obj.PropSet {
			fmt.Printf("%v (%v:%v)\n", prop.Val, obj.Obj.Type, obj.Obj.Value)
		}
	}
}
