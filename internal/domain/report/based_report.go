package report

import (
	"github.com/rock2z/tradebot/internal/domain/strategy"
	"github.com/rock2z/tradebot/internal/domain/tberr"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
)

type BasedReport struct {
	units []IReportUnit
}

func NewBasedReport(units []IReportUnit) *BasedReport {
	return &BasedReport{units: units}
}

func (b *BasedReport) Add(unit IReportUnit) {
	b.units = append(b.units, unit)
}

func (b BasedReport) GetReportUnit(slot timeslot.ISlot) (IReportUnit, error) {
	for _, unit := range b.units {
		if unit.GetSlot().Equal(slot) {
			return unit, nil
		}
	}
	return nil, tberr.ErrNotFound
}

func (b BasedReport) GetReportUnits() []IReportUnit {
	return b.units
}

type BasedReportUnit struct {
	slot      timeslot.ISlot
	operation strategy.Operation

	price  float64
	asset  float64
	cash   float64
	equity int64
}

func NewBasedReportUnit(
	slot timeslot.ISlot, operation strategy.Operation, price float64, asset float64, cash float64, equity int64,
) *BasedReportUnit {
	return &BasedReportUnit{
		slot:      slot,
		operation: operation,
		price:     price,
		asset:     asset,
		cash:      cash,
		equity:    equity,
	}
}

func (b BasedReportUnit) GetOperation() strategy.Operation {
	return b.operation
}

func (b BasedReportUnit) GetSlot() timeslot.ISlot {
	return b.slot
}

func (b BasedReportUnit) GetPrice() float64 {
	return b.price
}

func (b BasedReportUnit) GetAsset() float64 {
	return b.asset
}

func (b BasedReportUnit) GetCash() float64 {
	return b.cash
}

func (b BasedReportUnit) GetEquity() int64 {
	return b.equity
}
