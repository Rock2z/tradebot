package timeslot

type ISlot interface {
	GetTimeStamp() int64
	Equal(slot ISlot) bool
}

type ISeries interface {
	GetIndex(slot ISlot) (int, error)
	GetSlot(index int) (ISlot, error)
	GetSlots() []ISlot
}
