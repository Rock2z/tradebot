package main

import (
	"math/rand"
	"time"

	"github.com/piquette/finance-go/datetime"
	"github.com/rock2z/tradebot/internal/domain/backtester"
	"github.com/rock2z/tradebot/internal/domain/report"
	"github.com/rock2z/tradebot/internal/domain/stock"
	"github.com/rock2z/tradebot/internal/domain/strategy/grid"
	"github.com/rock2z/tradebot/internal/util"
	"go.uber.org/zap"
)

func main() {
	rand.Seed(time.Now().Unix())
	InitLogger()
	Run()
}

func Run() {
	// param definition
	symbol := "AAPL"
	startTime, _ := time.ParseInLocation(util.DefaultShortLayout, "2021-Jan-01", util.USLocation)
	endTime := time.Now().In(util.USLocation)
	interval := datetime.OneDay
	capital := float64(1_000_000)

	// init iStock
	yahooStock := stock.NewYahooStock(symbol, startTime, endTime, interval)
	err := yahooStock.Init()
	if err != nil {
		zap.S().Fatalf("init yahoo stock fail|err=%v", err)
	}

	// init time series from stock
	timeSeries := yahooStock.GetTimeSeries()

	// init strategy
	st := grid.NewGridStrategy(capital)

	//init report
	rep := report.NewBasedReport(make([]report.IReportUnit, 0, len(timeSeries.GetSlots())))

	b := &backtester.BackTester{
		Strategy:   st,
		Report:     rep,
		TimeSeries: timeSeries,
		Stock:      yahooStock,
		Cash:       capital,
		Equity:     0,
	}
	err = b.BackTest()
	if err != nil {
		zap.S().Fatalf("back test fail|err=%v", err)
	}
}
