package backtester

import (
	"time"

	"github.com/rock2z/tradebot/internal/domain/operation"
	"github.com/rock2z/tradebot/internal/domain/property"
	"github.com/rock2z/tradebot/internal/domain/report"
	"github.com/rock2z/tradebot/internal/domain/report/analysis"
	"github.com/rock2z/tradebot/internal/domain/stock"
	"github.com/rock2z/tradebot/internal/domain/strategy"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
	"github.com/rock2z/tradebot/internal/util"
	"go.uber.org/zap"
)

type BackTester struct {
	Strategy strategy.IStrategy
	Report   report.IReport

	TimeSeries timeslot.ISeries
	Stock      stock.IStock
	Cash       float64
	Equity     int64
}

func (b *BackTester) BackTest() error {
	series := b.TimeSeries
	for _, now := range series.GetSlots() {
		unit := b.Stock.GetUnit(now)
		price := b.Strategy.GetPrice(unit)

		op := b.Strategy.Decide(now, b.Stock)
		if err := Operate(op, b.Strategy.GetProperty(), b.Strategy.GetPrice(unit)); err != nil {
			return err
		}

		asset := b.Strategy.GetProperty().GetAsset(price)
		zap.S().Infof("time=%s, %s, price=%v, asset=%f, cash=%f, equity=%d",
			time.UnixMilli(now.GetTimeStamp()).Format(util.DefaultLogLayout), op, price, asset, b.Cash, b.Equity)
		b.Report.Add(report.NewBasedReportUnit(now, op.GetOperationType(), price, asset, b.Cash, b.Equity))
	}

	return GenerateChart(b.Report)
}

func GenerateChart(r report.IReport) error {
	err := util.Loop(
		func() error {
			return analysis.GenerateChart(r, analysis.CategoryPrice)
		},
		func() error {
			return analysis.GenerateChart(r, analysis.CategoryAsset)
		},
		func() error {
			return analysis.GenerateChart(r, analysis.CategoryCash)
		},
		func() error {
			return analysis.GenerateChart(r, analysis.CategoryEquity)
		},
	)
	if err != nil {
		zap.S().Warnf("GenerateChart fail|err=%v", err)
		return err
	}
	return nil
}

//TODO put price in operation
func Operate(op operation.IOperation, p property.IProperty, price float64) error {
	switch op.GetOperationType() {
	case operation.BUY:
		return p.Buy(price, op.GetAmount())
	case operation.SELL:
		return p.Sell(price, op.GetAmount())
	case operation.HOLD:
		fallthrough
	default:
		return nil
	}
}
