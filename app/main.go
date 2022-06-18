package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/rock2z/tradebot/internal/domain/backtester"
	"github.com/rock2z/tradebot/internal/domain/report"
	"github.com/rock2z/tradebot/internal/domain/stock"
	"github.com/rock2z/tradebot/internal/domain/strategy/buy_and_hold"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
	"github.com/rock2z/tradebot/internal/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	rand.Seed(time.Now().Unix())

	InitLogger()

	err := os.Setenv("TZ", "America/New_York")
	if err != nil {
		zap.S().Fatalf("set local time zone fail|err=%v", err)
	}
	defer func() {
		err = os.Unsetenv("TZ")
		if err != nil {
			zap.S().Fatalf("unset local time zone fail|err=%v", err)
		}
	}()

	Run()
}

func InitLogger() {
	level := zap.DebugLevel
	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = zapcore.TimeEncoderOfLayout(util.DefaultLayout)
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
	yahooStock := stock.NewYahooStock("AAPL", time.UnixMilli(1655284246_000), time.Now())
	err := yahooStock.Init()
	if err != nil {
		zap.S().Fatalf("init yahoo stock fail|err=%v", err)
	}

	units := yahooStock.GetUnits()
	series := make([]timeslot.ISlot, 0, len(units))
	for _, unit := range units {
		series = append(series, unit.GetSlot())
	}

	reportUnits := make([]report.IReportUnit, 0, len(units))

	timeSeries := timeslot.NewBasedSeries(series)
	b := &backtester.BackTester{
		Strategy:   buy_and_hold.NewBuyAndHoldStrategy(timeSeries, yahooStock),
		Report:     report.NewBasedReport(reportUnits),
		TimeSeries: timeSeries,
		Stock:      yahooStock,
		Cash:       1_000_000,
		Equity:     0,
	}
	err = b.BackTest()
	if err != nil {
		zap.S().Fatalf("back test fail|err=%v", err)
	}
}
