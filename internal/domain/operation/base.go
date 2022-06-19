package operation

type Operation struct {
	operationType Type
	amount        int64
}

func (o *Operation) GetOperationType() Type {
	return o.operationType
}

func (o *Operation) GetAmount() int64 {
	return o.amount
}

func NewBuyOperation(amount int64) *Operation {
	return &Operation{
		operationType: BUY,
		amount:        amount,
	}
}

func NewSellOperation(amount int64) *Operation {
	return &Operation{
		operationType: SELL,
		amount:        amount,
	}
}

func NewHoldOperation() *Operation {
	return &Operation{
		operationType: HOLD,
	}
}
