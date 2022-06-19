package grid

import (
	"math"
	"sync"

	"github.com/rock2z/tradebot/internal/domain/stock"
	"github.com/rock2z/tradebot/internal/domain/strategy"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
	"go.uber.org/zap"
)

type Strategy struct {
	once sync.Once

	base float64
	//even grid
	interval float64

	currentGridNumber int64
}

func (s *Strategy) Decide(slot timeslot.ISlot, stock stock.IStock) strategy.Operation {
	unit, err := stock.GetUnit(slot)
	if err != nil {
		zap.S().Warnf("stock.GetUnit(slot) fail|slot=%v|err=%v", slot, err)
		return strategy.Hold
	}
	s.once.Do(func() {
		s.Init(unit)
	})

	defer func() { s.currentGridNumber = s.GetGridNumber(unit) }()
	if s.GetGridNumber(unit) < s.currentGridNumber {
		return strategy.Buy
	}
	if s.GetGridNumber(unit) > s.currentGridNumber {
		return strategy.Sell
	}
	return strategy.Hold
}

func (s *Strategy) Init(unit stock.IStockUnit) {
	s.base = s.GetPrice(unit)
	s.interval = math.Abs(unit.GetHigh()-unit.GetLow()) / 2.0
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
