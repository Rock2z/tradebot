package timeslot

import "time"

type BasedSlot struct {
	// use time.Time
	timestamp time.Time
}

func NewBasedSlot(t time.Time) *BasedSlot {
	return &BasedSlot{timestamp: t}
}

func (b BasedSlot) GetTimeStamp() int64 {
	return b.timestamp.UnixMilli()
}

func (b BasedSlot) Equal(slot ISlot) bool {
	return b.GetTimeStamp() == slot.GetTimeStamp()
}
