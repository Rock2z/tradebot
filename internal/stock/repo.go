package stock

import (
	"github.com/rock2z/tradebot/pb"
)

type IRepo interface {
	GetLatestQuote(symbol string) *pb.Stock
	GetHistoricalQuote(symbol string, timestamp int64) *pb.Stock
}
