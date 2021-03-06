package impl

import (
	"math/rand"

	"github.com/rock2z/tradebot/internal/domain/stock"
	"github.com/rock2z/tradebot/internal/domain/strategy"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
)

type RandomStrategy struct{}

func (b RandomStrategy) Decide(time timeslot.ISlot, stock stock.IStock) strategy.Operation {
	return strategy.Operation(rand.Int63n(3))
}
