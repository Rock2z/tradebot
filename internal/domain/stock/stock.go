package stock

import "github.com/rock2z/tradebot/internal/domain/timeslot"

type IStock interface {
	GetTimeSeries() timeslot.ISeries
	GetUnits() []IStockUnit
	GetUnit(slot timeslot.ISlot) (IStockUnit, error)
}

type IStockUnit interface {
	GetSlot() timeslot.ISlot
	GetHigh() float64
	GetLow() float64
	GetOpen() float64
	GetClose() float64
	GetVolume() float64
}
