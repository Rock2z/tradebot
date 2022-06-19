package buy_and_hold

import (
	"sync"

	"github.com/rock2z/tradebot/internal/domain/stock"
	"github.com/rock2z/tradebot/internal/domain/strategy"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
)

type BuyAndHoldStrategy struct {
	once sync.Once
}

func (b *BuyAndHoldStrategy) Decide(slot timeslot.ISlot, stock stock.IStock) strategy.Operation {
	op := strategy.Hold
	b.once.Do(func() {
		op = strategy.Buy
	})
	return op
}
