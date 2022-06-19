package grid

import (
	"math"
	"sync"

	"github.com/rock2z/tradebot/internal/domain/operation"
	"github.com/rock2z/tradebot/internal/domain/property"
	"github.com/rock2z/tradebot/internal/domain/stock"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
	"github.com/rock2z/tradebot/internal/util"
)

type Strategy struct {
	once sync.Once

	property property.IProperty

	//even grid
	base              float64
	interval          float64
	currentGridNumber int64
}

func NewGridStrategy(capital float64) *Strategy {
	return &Strategy{
		property: property.NewProperty(capital, 0),
	}
}

func (s *Strategy) Decide(slot timeslot.ISlot, stock stock.IStock) operation.IOperation {
	unit := stock.GetUnit(slot)
	s.once.Do(func() {
		s.Init(unit)
	})

	defer func() { s.currentGridNumber = s.GetGridNumber(unit) }()
	if s.GetGridNumber(unit) < s.currentGridNumber {
		return operation.NewBuyOperation(util.CalcMaxShare(s.GetProperty().GetCash(), s.GetPrice(unit)))
	}
	if s.GetGridNumber(unit) > s.currentGridNumber {
		return operation.NewSellOperation(s.GetProperty().GetEquity())
	}
	return operation.NewHoldOperation()
}

func (s *Strategy) GetProperty() property.IProperty {
	return s.property
}

func (s *Strategy) Init(unit stock.IStockUnit) {
	s.base = s.GetPrice(unit)
	s.interval = math.Abs(unit.GetHigh()-unit.GetLow()) / float64(2)
	s.currentGridNumber = s.GetGridNumber(unit)
}

func (s *Strategy) GetPrice(unit stock.IStockUnit) float64 {
	return unit.GetClose()
}

func (s *Strategy) GetGridNumber(unit stock.IStockUnit) int64 {
	diff := s.GetPrice(unit) - s.base
	num := diff / s.interval

	//think: what if I return float here?
	return int64(math.Trunc(num))
}
