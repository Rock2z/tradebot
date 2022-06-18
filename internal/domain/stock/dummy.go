package stock

import (
	"math/rand"

	"github.com/rock2z/tradebot/internal/domain/timeslot"
	"go.uber.org/zap"
)

type DummyStock struct {
	Val float64
}

func (d DummyStock) GetHigh(time timeslot.ISlot) float64 {
	panic("implement me")
}

func (d DummyStock) GetLow(time timeslot.ISlot) float64 {
	panic("implement me")
}

func (d DummyStock) GetOpen(time timeslot.ISlot) float64 {
	diff := rand.Int63n(100)
	op := rand.Int63n(2) == 0
	if !op && d.Val-float64(diff) > 0 {
		diff -= 2 * diff
	}
	val := d.Val + float64(diff)
	zap.S().Infof("current price=%f, diff=%d", val, diff)
	return val
}

func (d DummyStock) GetClose(time timeslot.ISlot) float64 {
	panic("implement me")
}

func (d DummyStock) GetVolume(time timeslot.ISlot) float64 {
	panic("implement me")
}
