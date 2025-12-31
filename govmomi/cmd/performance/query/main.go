package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vmware/govmomi/performance"
	"github.com/vmware/govmomi/vim25/types"

	sx "github.com/9506hqwy/samples-go/pkg/sessionex"
)

func main() {
	ctx := context.Background()

	url := os.Getenv("VSPHERE_URL")
	user := os.Getenv("VSPHERE_USER")
	pass := os.Getenv("VSPHERE_PASSWORD")

	// format = motype:value
	moref := strings.Split(os.Args[1], ":")

	// format = counterid:instance
	metric := strings.Split(os.Args[2], ":")
	num, err := strconv.ParseInt(metric[0], 10, 32)
	if err != nil {
		panic(err)
	}

	c, err := sx.Login(ctx, url, user, pass)
	if err != nil {
		panic(err)
	}

	defer sx.Logout(ctx, c)

	pm := performance.NewManager(c)

	mor := types.ManagedObjectReference{
		Type:  moref[0],
		Value: moref[1],
	}
	metricId := types.PerfMetricId{
		CounterId: int32(num),
		Instance:  metric[1],
	}
	spec := types.PerfQuerySpec{
		Entity:   mor,
		MetricId: []types.PerfMetricId{metricId},
	}
	series, err := pm.Query(ctx, []types.PerfQuerySpec{spec})
	if err != nil {
		panic(err)
	}

	metrics, err := pm.ToMetricSeries(ctx, series)
	if err != nil {
		panic(err)
	}

	for _, metric := range metrics {
		for _, value := range metric.Value {
			for idx, v := range value.Value {
				sample := metric.SampleInfo[idx]

				fmt.Printf("%v %v\n", sample.Timestamp, v)
			}
		}
	}
}
