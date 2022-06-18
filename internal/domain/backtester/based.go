package backtester

import (
	"math"

	"github.com/rock2z/tradebot/internal/domain/stock"
	"github.com/rock2z/tradebot/internal/domain/strategy"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
	"github.com/rock2z/tradebot/internal/util/chart"
	"go.uber.org/zap"
)

type BackTester struct {
	Strategy strategy.IStrategy

	TimeSeries timeslot.ISeries
	Stock      stock.IStock
	Cash       float64
	Equity     int64
}

func (b *BackTester) BackTest() error {
	assetArr := make([]float64, 0)
	priceArr := make([]float64, 0)

	series := b.TimeSeries
	for {
		now := series.GetCurrent()
		if !series.HasMore() {
			break
		}

		price, err := b.Stock.GetClose(now)
		if err != nil {
			zap.S().Infof("b.Stock.GetOpen fail, current slot = %v", series.GetCurrent())
			return err
		}

		op := b.Strategy.Decide(now)
		o := "HOLD"
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
			o = "BUY"
		case strategy.Sell:
			if b.Equity <= 0 {
				zap.S().Infof("want SELL, but have no stock, so HOLD")
				break
			}
			b.Cash += float64(b.Equity) * price
			b.Equity = 0
			o = "SELL"
		case strategy.Hold:
			fallthrough
		default:
		}
		asset := float64(b.Equity)*price + b.Cash
		zap.S().Infof("%v, current asset=%f, cash=%f, equity=%d", o, asset, b.Cash, b.Equity)
		assetArr = append(assetArr, asset)
		priceArr = append(priceArr, price)

		err = series.Next()
		if err != nil {
			zap.S().Infof("Next fail, current slot = %v", series.GetCurrent())
			return err
		}
	}
	chart.CreateLineChart("assetArr", assetArr)
	chart.CreateLineChart("priceArr", priceArr)
	return nil
}
