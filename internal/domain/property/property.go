package property

type IProperty interface {
	GetCash() float64
	GetEquity() int64
	GetAsset(price float64) float64
	Buy(price float64, amount int64) error
	Sell(price float64, amount int64) error
}
