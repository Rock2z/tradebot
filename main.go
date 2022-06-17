package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/rock2z/tradebot/internal"
)

type IStock interface {
	GetHigh(time ITimeSlot) float64
	GetLow(time ITimeSlot) float64
	GetOpen(time ITimeSlot) float64
	GetClose(time ITimeSlot) float64
	GetVolume(time ITimeSlot) float64
}

/*
IStrategy is used to define how this strategy work.
To simulate a complex strategy, we can divide the capital money to several parts. We can decide each part own how much money.

Operation function's result means, what operation we want to apply to this part of money, on a certain date.
When a strategy run, Operation function will be called every timeslot from the startTime to endTime.
*/
type IStrategy interface {
	Decide(time ITimeSlot, stock IStock) Operation
}

type Operation uint32

const (
	Buy Operation = iota
	Sell
	Hold
)

type ITimeSlot interface {
	GetTimeStamp() int64
	GetYear() int64
	GetMonth() int64
	GetDay() int64
	GetSlot() int64
}

type BasedTimeSlot struct {
	timestamp int64
}

func NewBasedTimeSlot(t int64) ITimeSlot {
	return &BasedTimeSlot{timestamp: t}
}

func (b BasedTimeSlot) GetTimeStamp() int64 {
	return b.timestamp
}

func (b BasedTimeSlot) GetYear() int64 {
	return int64(time.UnixMilli(b.timestamp).Year())
}

func (b BasedTimeSlot) GetMonth() int64 {
	return int64(time.UnixMilli(b.timestamp).Month())
}

func (b BasedTimeSlot) GetDay() int64 {
	return int64(time.UnixMilli(b.timestamp).Day())
}

func (b BasedTimeSlot) GetSlot() int64 {
	return int64(time.UnixMilli(b.timestamp).Sub(time.UnixMilli(0)).Hours() / 24)
}

/*
IReport asset = equity + cash
*/
type IReport interface {
	GetAsset(time ITimeSlot) float64
	GetEquity(time ITimeSlot) int64
	GetCash(time ITimeSlot) float64
}

type BasedStrategy struct{}

func (b BasedStrategy) Decide(time ITimeSlot, stock IStock) Operation {
	return Operation(rand.Int63n(3))
}

type DummyStock struct {
	val float64
}

func (d DummyStock) GetHigh(time ITimeSlot) float64 {
	//TODO implement me
	panic("implement me")
}

func (d DummyStock) GetLow(time ITimeSlot) float64 {
	//TODO implement me
	panic("implement me")
}

func (d DummyStock) GetOpen(time ITimeSlot) float64 {
	diff := rand.Int63n(100)
	op := rand.Int63n(2) == 0
	if !op && d.val-float64(diff) > 0 {
		diff -= 2 * diff
	}
	val := d.val + float64(diff)
	log.Printf("current price=%f, diff=%d", val, diff)
	return val
}

func (d DummyStock) GetClose(time ITimeSlot) float64 {
	//TODO implement me
	panic("implement me")
}

func (d DummyStock) GetVolume(time ITimeSlot) float64 {
	//TODO implement me
	panic("implement me")
}

type BackTester struct {
	strategy IStrategy

	startTime, endTime ITimeSlot
	stock              IStock
	cash               float64
	equity             int64
}

func (b *BackTester) Init() {
	b.strategy = &BasedStrategy{}
	b.startTime = NewBasedTimeSlot(1650188245000)
	b.endTime = NewBasedTimeSlot(time.Now().UnixMilli())
	b.stock = &DummyStock{val: 1000}
	b.cash = 1_000_000
	b.equity = 0
}

func (b *BackTester) BackTest() {
	assetArr := make([]float64, 0)
	priceArr := make([]float64, 0)
	for i := b.startTime.GetSlot(); i <= b.endTime.GetSlot(); i++ {
		now := NewBasedTimeSlot(b.startTime.GetTimeStamp() + (i * 24 * 60 * 60 * 1000))
		price := b.stock.GetOpen(now)

		op := b.strategy.Decide(now, b.stock)
		o := "HOLD"
		switch op {
		case Buy:
			share := int64(math.Trunc(b.cash / price))
			cost := float64(share) * price
			if b.cash < cost {
				log.Printf("want BUY, but poor, so HOLD")
				break
			}
			b.equity += share
			b.cash -= cost
			o = "BUY"
		case Sell:
			if b.equity <= 0 {
				log.Printf("want SELL, but have no stock, so HOLD")
				break
			}
			b.cash += float64(b.equity) * price
			b.equity = 0
			o = "SELL"
		case Hold:
			fallthrough
		default:
		}
		asset := float64(b.equity)*price + b.cash
		log.Printf("%v, current asset=%f, cash=%f, equity=%d\n\n", o, asset, b.cash, b.equity)
		assetArr = append(assetArr, asset)
		priceArr = append(priceArr, price)
	}
	internal.CreateLineChart("assetArr", assetArr)
	internal.CreateLineChart("priceArr", priceArr)
}

func main() {
	rand.Seed(time.Now().Unix())
	b := &BackTester{
		strategy:  &BasedStrategy{},
		startTime: NewBasedTimeSlot(1650188245000),
		endTime:   NewBasedTimeSlot(time.Now().UnixMilli()),
		stock:     &DummyStock{val: 1000},
		cash:      1_000_000,
		equity:    0,
	}
	b.BackTest()
}
