package property

import (
	"github.com/rock2z/tradebot/internal/domain/tberr"
)

type Property struct {
	cash   float64
	equity int64
}

func NewProperty(cash float64, equity int64) *Property {
	return &Property{
		cash:   cash,
		equity: equity,
	}
}

func (p *Property) GetCash() float64 {
	return p.cash
}

func (p *Property) GetEquity() int64 {
	return p.equity
}

func (p *Property) GetAsset(price float64) float64 {
	return float64(p.equity)*price + p.cash
}

func (p *Property) Buy(price float64, amount int64) error {
	if price <= 0 || amount < 0 {
		return tberr.ErrInvalidParam
	}

	cost := price * float64(amount)
	if cost > p.cash {
		return tberr.ErrCashNotEnough
	}
	p.equity += amount
	p.cash -= cost
	return nil
}

func (p *Property) Sell(price float64, amount int64) error {
	if price <= 0 || amount < 0 {
		return tberr.ErrInvalidParam
	}
	if amount > p.equity {
		return tberr.ErrNoEnoughShare
	}

	profit := price * float64(amount)
	p.equity -= amount
	p.cash += profit
	return nil
}
