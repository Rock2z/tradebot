package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/rock2z/tradebot/internal/domain/backtester"
	"github.com/rock2z/tradebot/internal/domain/stock"
	"github.com/rock2z/tradebot/internal/domain/strategy/impl"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
)

func main() {
	rand.Seed(time.Now().Unix())
	yahooStock := stock.NewYahooStock("AAPL", time.UnixMilli(1623942000_000), time.Now())
	err := yahooStock.Init()
	if err != nil {
		log.Fatalf("init yahoo stock fail|err=%v", err)
	}

	units := yahooStock.GetUnits()
	series := make([]timeslot.ISlot, 0, len(units))
	for _, unit := range units {
		series = append(series, unit.GetSlot())
	}

	b := &backtester.BackTester{
		Strategy:   &impl.RandomStrategy{},
		TimeSeries: timeslot.NewBasedSeries(series),
		Stock:      yahooStock,
		Cash:       1_000_000,
		Equity:     0,
	}
	err = b.BackTest()
	if err != nil {
		log.Fatalf("back test fail|err=%v", err)
	}
}
