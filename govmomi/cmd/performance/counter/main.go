package main

import (
	"context"
	"fmt"
	"os"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"

	px "github.com/9506hqwy/samples-go/pkg/performanceex"
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

	var p mo.PerformanceManager
	err = pc.RetrieveOne(ctx, *c.ServiceContent.PerfManager, nil, &p)
	if err != nil {
		panic(err)
	}

	for _, counter := range p.PerfCounter {
		label := counter.GroupInfo.GetElementDescription().Label
		name := counter.NameInfo.GetElementDescription().Label
		lvl := counter.Level
		devicelvl := counter.PerDeviceLevel
		rollup := px.GetCounterTypeLabel(&p, counter.RollupType)
		stat := px.GetStatTypeLabel(&p, counter.StatsType)
		unit := counter.UnitInfo.GetElementDescription().Label
		fmt.Printf("%v Group:%v Name:%v Level:%v DeviceLevel:%v Rollup:%v Stat:%v Unit:%v\n", counter.Key, label, name, lvl, devicelvl, rollup, stat, unit)
	}
}
