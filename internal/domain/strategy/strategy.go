package strategy

import (
	"github.com/rock2z/tradebot/internal/domain/stock"
	"github.com/rock2z/tradebot/internal/domain/timeslot"
)

type Operation uint32

const (
	Buy Operation = iota
	Sell
	Hold
)

/*
IStrategy is used to define how this strategy work.
To simulate a complex strategy, we can divide the capital money to several parts. We can decide each part own how much money.

Operation function's result means, what operation we want to apply to this part of money, on a certain date.
When a strategy run, Operation function will be called every timeslot from the startTime to endTime.
*/
type IStrategy interface {
	Decide(time timeslot.ISlot, stock stock.IStock) Operation
}
