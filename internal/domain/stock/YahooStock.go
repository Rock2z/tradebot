package stock

import (
	"fmt"
	"time"

	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/rock2z/tradebot/internal/domain/tberr"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
)

type YahooStock struct {
	symbol   string
	from, to time.Time
	units    []IStockUnit
}

func NewYahooStock(symbol string, from, to time.Time) *YahooStock {
	return &YahooStock{
		symbol: symbol,
		from:   from,
		to:     to,
		units:  make([]IStockUnit, 0),
	}
}

func (y *YahooStock) Init() error {
	params := &chart.Params{
		Symbol:   y.symbol,
		Start:    datetime.New(&y.from),
		End:      datetime.New(&y.to),
		Interval: datetime.OneMin,
	}
	iter := chart.Get(params)
	for iter.Next() {
		bar := iter.Current().(*finance.ChartBar)
		//if !util.InRegularMarketingTime(time.Unix(int64(bar.Timestamp), 0)) {
		//	continue
		//}
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

func (y *YahooStock) GetUnits() []IStockUnit {
	return y.units
}

func (y *YahooStock) GetHigh(slot timeslot.ISlot) (float64, error) {
	for _, unit := range y.units {
		if unit.GetSlot().Equal(slot) {
			return unit.GetHigh(), nil
		}
	}
	return 0, tberr.ErrNotFound
}

func (y *YahooStock) GetLow(slot timeslot.ISlot) (float64, error) {
	for _, unit := range y.units {
		if unit.GetSlot().Equal(slot) {
			return unit.GetLow(), nil
		}
	}
	return 0, tberr.ErrNotFound
}

func (y *YahooStock) GetOpen(slot timeslot.ISlot) (float64, error) {
	for _, unit := range y.units {
		if unit.GetSlot().Equal(slot) {
			return unit.GetOpen(), nil
		}
	}
	return 0, tberr.ErrNotFound
}

func (y *YahooStock) GetClose(slot timeslot.ISlot) (float64, error) {
	for _, unit := range y.units {
		if unit.GetSlot().Equal(slot) {
			return unit.GetClose(), nil
		}
	}
	return 0, tberr.ErrNotFound
}

func (y *YahooStock) GetVolume(slot timeslot.ISlot) (float64, error) {
	for _, unit := range y.units {
		if unit.GetSlot().Equal(slot) {
			return unit.GetVolume(), nil
		}
	}
	return 0, tberr.ErrNotFound
}
