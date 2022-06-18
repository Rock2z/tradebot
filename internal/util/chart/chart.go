package chart

import (
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func CreateLineChart(name string, data []float64) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    name,
			Subtitle: "One Year Performance",
		}),
	)
	// Put data into instance
	items := make([]opts.LineData, 0)
	for _, e := range data {
		items = append(items, opts.LineData{Value: e})
	}
	line.SetXAxis(getXAxis(data)).
		AddSeries(name, items).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	f, _ := os.Create("log/" + name + ".html")
	_ = line.Render(f)
}

func getXAxis(a []float64) []string {
	ret := make([]string, 0, len(a))
	for i := range a {
		ret = append(ret, strconv.Itoa(i))
	}
	return ret
}
