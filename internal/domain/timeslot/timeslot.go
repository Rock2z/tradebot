package timeslot

type ISlot interface {
	GetTimeStamp() int64
	Equal(slot ISlot) bool
}

type ISeries interface {
	GetCurrent() ISlot
	GetSlot(index int) (ISlot, error)
	GetIndex(slot ISlot) (int, error)
	Next() error
	HasMore() bool
	Reset()
}
