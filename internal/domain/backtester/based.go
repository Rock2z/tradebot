package backtester

import (
	"math"
	"time"

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
		price, err := b.Stock.GetClose(now)
		if err != nil {
			zap.S().Infof("b.Stock.GetOpen fail, current slot = %v", now)
			return err
		}

		op := b.Strategy.Decide(now)
		switch op {
		case strategy.Buy:
			share := int64(math.Trunc(b.Cash / price))
			cost := float64(share) * price
			if b.Cash < cost {
				zap.S().Infof("want BUY, but poor, so HOLD")
				break
			}
			b.Equity += share
			b.Cash -= cost
		case strategy.Sell:
			if b.Equity <= 0 {
				zap.S().Infof("want SELL, but have no stock, so HOLD")
				break
			}
			b.Cash += float64(b.Equity) * price
			b.Equity = 0
		case strategy.Hold:
			fallthrough
		default:
		}
		asset := float64(b.Equity)*price + b.Cash
		zap.S().Infof("time=%s, %s, price=%v, asset=%f, cash=%f, equity=%d",
			time.UnixMilli(now.GetTimeStamp()).Format(util.DefaultLogLayout), op, price, asset, b.Cash, b.Equity)
		b.Report.Add(report.NewBasedReportUnit(now, op, price, asset, b.Cash, b.Equity))
	}
	err := util.Loop(
		func() error {
			return analysis.GenerateChart(b.Report, analysis.CategoryPrice)
		},
		func() error {
			return analysis.GenerateChart(b.Report, analysis.CategoryAsset)
		},
		func() error {
			return analysis.GenerateChart(b.Report, analysis.CategoryCash)
		},
		func() error {
			return analysis.GenerateChart(b.Report, analysis.CategoryEquity)
		},
	)
	if err != nil {
		zap.S().Warnf("GenerateChart fail|err=%v", err)
		return err
	}
	return nil
}
