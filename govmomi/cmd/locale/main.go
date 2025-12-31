package main

import (
	"context"
	"fmt"
	"os"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"

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

	pc := property.DefaultCollector(c)

	var s mo.SessionManager
	err = pc.RetrieveOne(ctx, *c.ServiceContent.SessionManager, []string{"supportedLocaleList"}, &s)
	if err != nil {
		panic(err)
	}

	for _, locale := range s.SupportedLocaleList {
		fmt.Printf("%v\n", locale)
	}
}
