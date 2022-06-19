package stock

import "github.com/rock2z/tradebot/internal/domain/timeslot"

type IStock interface {
	GetTimeSeries() timeslot.ISeries
	GetUnits() []IStockUnit
	GetHigh(time timeslot.ISlot) (float64, error)
	GetLow(time timeslot.ISlot) (float64, error)
	GetOpen(time timeslot.ISlot) (float64, error)
	GetClose(time timeslot.ISlot) (float64, error)
	GetVolume(time timeslot.ISlot) (float64, error)
}

type IStockUnit interface {
	GetSlot() timeslot.ISlot
	GetHigh() float64
	GetLow() float64
	GetOpen() float64
	GetClose() float64
	GetVolume() float64
}
