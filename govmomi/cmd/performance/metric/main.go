package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/vmware/govmomi/performance"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"

	px "github.com/9506hqwy/samples-go/pkg/performanceex"
	sx "github.com/9506hqwy/samples-go/pkg/sessionex"
)

func main() {
	ctx := context.Background()

	url := os.Getenv("VSPHERE_URL")
	user := os.Getenv("VSPHERE_USER")
	pass := os.Getenv("VSPHERE_PASSWORD")

	// format = motype:value
	moref := strings.Split(os.Args[1], ":")

	c, err := sx.Login(ctx, url, user, pass)
	if err != nil {
		panic(err)
	}

	defer sx.Logout(ctx, c)

	pc := property.DefaultCollector(c)

	var p mo.PerformanceManager
	err = pc.RetrieveOne(ctx, *c.ServiceContent.PerfManager, nil, &p)
	if err != nil {
		panic(err)
	}

	pm := performance.NewManager(c)

	mor := types.ManagedObjectReference{
		Type:  moref[0],
		Value: moref[1],
	}
	metrics, err := pm.AvailableMetric(ctx, mor, 0)
	if err != nil {
		panic(err)
	}

	for _, metric := range metrics {
		counter := px.GetCounter(&p, metric.CounterId)

		name := ""
		if counter != nil {
			name = counter.NameInfo.GetElementDescription().Label
		}

		fmt.Printf("%v(%v:%v)\n", name, counter.Key, metric.Instance)
	}
}
