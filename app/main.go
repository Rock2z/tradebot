package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/piquette/finance-go/datetime"
	"github.com/rock2z/tradebot/internal/domain/backtester"
	"github.com/rock2z/tradebot/internal/domain/report"
	"github.com/rock2z/tradebot/internal/domain/stock"
	"github.com/rock2z/tradebot/internal/domain/strategy/buy_and_hold"
	"github.com/rock2z/tradebot/internal/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	rand.Seed(time.Now().Unix())
	InitLogger()

	Run()
}

func InitLogger() {
	level := zap.DebugLevel
	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = zapcore.TimeEncoderOfLayout(util.DefaultLogLayout)
	fileEncoder := zapcore.NewJSONEncoder(pe)
	consoleEncoder := zapcore.NewConsoleEncoder(pe)
	f, _ := os.Create("log/data.log")

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(f), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

func Run() {
	// param definition
	symbol := "AAPL"
	startTime, _ := time.ParseInLocation(util.DefaultShortLayout, "2022-Jun-14", util.USLocation)
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
	strategy := buy_and_hold.NewBuyAndHoldStrategy(timeSeries, yahooStock)

	//init report
	rep := report.NewBasedReport(make([]report.IReportUnit, 0, len(timeSeries.GetSlots())))

	b := &backtester.BackTester{
		Strategy:   strategy,
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
