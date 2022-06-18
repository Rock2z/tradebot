package stock

import "github.com/rock2z/tradebot/internal/domain/timeslot"

type BasedStockUnit struct {
	slot                           timeslot.ISlot
	high, low, open, close, volume float64
}

func NewBasedStockUnit(slot timeslot.ISlot, high, low, open, close, volume float64) *BasedStockUnit {
	return &BasedStockUnit{
		slot:   slot,
		high:   high,
		low:    low,
		open:   open,
		close:  close,
		volume: volume,
	}
}

func (b BasedStockUnit) GetSlot() timeslot.ISlot {
	return b.slot
}

func (b BasedStockUnit) GetHigh() float64 {
	return b.high
}

func (b BasedStockUnit) GetLow() float64 {
	return b.low
}

func (b BasedStockUnit) GetOpen() float64 {
	return b.open
}

func (b BasedStockUnit) GetClose() float64 {
	return b.close
}

func (b BasedStockUnit) GetVolume() float64 {
	return b.volume
}
