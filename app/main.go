package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/rock2z/tradebot/internal/domain/backtester"
	"github.com/rock2z/tradebot/internal/domain/stock"
	"github.com/rock2z/tradebot/internal/domain/strategy/buy_and_hold"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	InitLogger()
	rand.Seed(time.Now().Unix())

	Run()
}

func InitLogger() {
	level := zap.DebugLevel
	pe := zap.NewProductionEncoderConfig()
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
	yahooStock := stock.NewYahooStock("AAPL", time.UnixMilli(1623942000_000), time.Now())
	err := yahooStock.Init()
	if err != nil {
		zap.S().Fatalf("init yahoo stock fail|err=%v", err)
	}

	units := yahooStock.GetUnits()
	series := make([]timeslot.ISlot, 0, len(units))
	for _, unit := range units {
		series = append(series, unit.GetSlot())
	}

	timeSeries := timeslot.NewBasedSeries(series)
	b := &backtester.BackTester{
		Strategy:   buy_and_hold.NewBuyAndHoldStrategy(timeSeries, yahooStock),
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
