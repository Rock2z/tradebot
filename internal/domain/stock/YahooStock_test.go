package stock

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestYahooStock_GetHigh(t *testing.T) {
	yahooStock := NewYahooStock("AAPL", time.Unix(0, time.Now().UnixNano()-int64(time.Hour*24*7)), time.Now())
	err := yahooStock.Init()
	assert.NoError(t, err)
	zap.S().Info(yahooStock)
}
