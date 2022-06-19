package buy_and_hold

import (
	"sync"

	"github.com/rock2z/tradebot/internal/domain/operation"
	"github.com/rock2z/tradebot/internal/domain/property"
	"github.com/rock2z/tradebot/internal/domain/stock"
	"github.com/rock2z/tradebot/internal/domain/tberr"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
	"github.com/rock2z/tradebot/internal/util"
)

type BuyAndHoldStrategy struct {
	once sync.Once

	property property.IProperty
}

func (b *BuyAndHoldStrategy) GetProperty() property.IProperty {
	return b.property
}

func (b *BuyAndHoldStrategy) Decide(slot timeslot.ISlot, stock stock.IStock) operation.IOperation {
	op := operation.NewHoldOperation()
	unit := stock.GetUnit(slot)
	b.once.Do(func() {
		op = operation.NewBuyOperation(util.CalcMaxShare(b.GetProperty().GetCash(), b.GetPrice(unit)))
	})
	return op
}

func (b *BuyAndHoldStrategy) GetPrice(unit stock.IStockUnit) float64 {
	return unit.GetClose()
}

func (b *BuyAndHoldStrategy) Operate(slot timeslot.ISlot, stock stock.IStock, op operation.IOperation) error {
	unit := stock.GetUnit(slot)
	switch op.GetOperationType() {
	case operation.BUY:
		return b.GetProperty().Buy(b.GetPrice(unit), op.GetAmount())
	case operation.SELL:
		return b.GetProperty().Sell(b.GetPrice(unit), op.GetAmount())
	case operation.HOLD:
		fallthrough
	default:
	}
	return tberr.ErrInvalidOperation
}
