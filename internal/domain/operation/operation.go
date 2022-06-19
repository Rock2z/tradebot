package operation

type Type string

const (
	BUY  Type = "BUY"
	SELL Type = "SELL"
	HOLD Type = "HOLD"
)

type IOperation interface {
	GetOperationType() Type
	GetAmount() int64
}
