package buy_and_hold

import (
	"github.com/rock2z/tradebot/internal/domain/stock"
	"github.com/rock2z/tradebot/internal/domain/strategy"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
	"go.uber.org/zap"
)

type BuyAndHoldStrategy struct {
	series timeslot.ISeries
	stock  stock.IStock
}

func NewBuyAndHoldStrategy(series timeslot.ISeries, stock stock.IStock) *BuyAndHoldStrategy {
	return &BuyAndHoldStrategy{
		series: series,
		stock:  stock,
	}
}

func (b *BuyAndHoldStrategy) Decide(slot timeslot.ISlot) strategy.Operation {
	firstSlot, err := b.series.GetSlot(0)
	if err != nil {
		zap.S().Warnf("fail to get first slot|series=%v", b.series)
		return strategy.Hold
	}
	if slot.Equal(firstSlot) {
		return strategy.Buy
	}
	return strategy.Hold
}
