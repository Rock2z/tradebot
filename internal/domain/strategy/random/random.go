package impl

import (
	"math/rand"

	"github.com/rock2z/tradebot/internal/domain/strategy"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
)

type RandomStrategy struct{}

func (b RandomStrategy) Decide(time timeslot.ISlot) strategy.Operation {
	switch rand.Int63n(3) {
	case 0:
		return strategy.Buy
	case 1:
		return strategy.Sell
	default:
		return strategy.Hold
	}
}
