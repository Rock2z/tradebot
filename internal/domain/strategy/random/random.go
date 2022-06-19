package impl

import (
	"math/rand"

	"github.com/rock2z/tradebot/internal/domain/operation"
	"github.com/rock2z/tradebot/internal/domain/property"
	"github.com/rock2z/tradebot/internal/domain/stock"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
	"github.com/rock2z/tradebot/internal/util"
)

type RandomStrategy struct {
	property property.IProperty
}

func (b *RandomStrategy) GetProperty() property.IProperty {
	return b.property
}

func (b RandomStrategy) Decide(slot timeslot.ISlot, stock stock.IStock) operation.IOperation {
	unit := stock.GetUnit(slot)
	switch rand.Int63n(3) {
	case 0:
		return operation.NewBuyOperation(util.CalcMaxShare(b.GetProperty().GetCash(), b.GetPrice(unit)))
	case 1:
		return operation.NewSellOperation(b.GetProperty().GetEquity())
	default:
		return operation.NewHoldOperation()
	}
}

func (b RandomStrategy) GetPrice(stock stock.IStockUnit) float64 {
	return stock.GetClose()
}
