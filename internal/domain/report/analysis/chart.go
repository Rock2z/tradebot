package analysis

import (
	"fmt"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/rock2z/tradebot/internal/domain/report"
	"github.com/rock2z/tradebot/internal/util"
)

type Category string

const (
	CategoryPrice  Category = "price"
	CategoryAsset  Category = "asset"
	CategoryCash   Category = "cash"
	CategoryEquity Category = "equity"
)

func GenerateChart(report report.IReport, category Category) error {
	units := report.GetReportUnits()
	XAxis := make([]string, 0, len(units))
	YAxis := make([]opts.LineData, 0, len(units))
	for _, unit := range units {
		slot := time.UnixMilli(unit.GetSlot().GetTimeStamp()).Format(util.DefaultLayout)
		XAxis = append(XAxis, slot)

		var yData float64
		switch category {
		case CategoryPrice:
			yData = unit.GetPrice()
		case CategoryAsset:
			yData = unit.GetAsset()
		case CategoryCash:
			yData = unit.GetCash()
		case CategoryEquity:
			yData = float64(unit.GetEquity())
		default:
		}
		YAxis = append(YAxis, opts.LineData{Value: yData})
	}

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: string(category),
		}),
	)

	line.SetXAxis(XAxis).
		AddSeries(string(category), YAxis).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	folder := fmt.Sprintf("report/%d/", time.Now().Unix())
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err = os.Mkdir(folder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
	f, err := os.Create(folder + string(category) + ".html")
	if err != nil {
		return err
	}
	return line.Render(f)
}
