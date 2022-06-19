package stock

import (
	"fmt"
	"time"

	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
)

type YahooStock struct {
	symbol   string
	from, to time.Time
	interval datetime.Interval
	units    []IStockUnit
}

func NewYahooStock(symbol string, from, to time.Time, interval datetime.Interval) *YahooStock {
	return &YahooStock{
		symbol:   symbol,
		from:     from,
		to:       to,
		interval: interval,
		units:    make([]IStockUnit, 0),
	}
}

func (y *YahooStock) Init() error {
	params := &chart.Params{
		Symbol:   y.symbol,
		Interval: y.interval,
		Start:    datetime.New(&y.from),
		End:      datetime.New(&y.to),
	}
	iter := chart.Get(params)
	for iter.Next() {
		bar := iter.Current().(*finance.ChartBar)
		unit := NewBasedStockUnit(
			timeslot.NewBasedSlot(time.Unix(int64(bar.Timestamp), 0)),
			bar.High.InexactFloat64(),
			bar.Low.InexactFloat64(),
			bar.Open.InexactFloat64(),
			bar.Close.InexactFloat64(),
			float64(bar.Volume),
		)
		y.units = append(y.units, unit)
	}
	if err := iter.Err(); err != nil {
		fmt.Println(err)
	}
	return nil
}

func (y *YahooStock) GetTimeSeries() timeslot.ISeries {
	units := y.GetUnits()
	series := make([]timeslot.ISlot, 0, len(units))
	for _, unit := range units {
		series = append(series, unit.GetSlot())
	}
	return timeslot.NewBasedSeries(series)
}

func (y *YahooStock) GetUnits() []IStockUnit {
	return y.units
}

func (y *YahooStock) GetUnit(slot timeslot.ISlot) IStockUnit {
	for _, unit := range y.units {
		if unit.GetSlot().Equal(slot) {
			return unit
		}
	}
	return nil
}
