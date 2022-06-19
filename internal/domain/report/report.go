package report

import (
	"github.com/rock2z/tradebot/internal/domain/operation"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
)

/*
IReport asset = equity + cash
*/
type IReport interface {
	Add(IReportUnit)
	GetReportUnit(timeslot.ISlot) (IReportUnit, error)
	GetReportUnits() []IReportUnit
}

type IReportUnit interface {
	GetSlot() timeslot.ISlot
	GetOperation() operation.Type
	GetPrice() float64
	GetAsset() float64
	GetCash() float64
	GetEquity() int64
}
