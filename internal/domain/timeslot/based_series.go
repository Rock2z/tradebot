package timeslot

import (
	"github.com/rock2z/tradebot/internal/domain/tberr"
)

type BasedSeries struct {
	slots []ISlot

	cur int
}

func NewBasedSeries(s []ISlot) *BasedSeries {
	return &BasedSeries{
		slots: s,
		cur:   0,
	}
}

func (b *BasedSeries) GetCurrent() ISlot {
	return b.slots[b.cur]
}

func (b *BasedSeries) GetSlots() []ISlot {
	return b.slots
}

func (b *BasedSeries) GetSlot(index int) (ISlot, error) {
	if index < 0 || index >= len(b.slots) {
		return nil, tberr.ErrNotFound
	}
	return b.slots[index], nil
}

func (b *BasedSeries) GetIndex(slot ISlot) (int, error) {
	for i, e := range b.slots {
		if slot.Equal(e) {
			return i, nil
		}
	}
	return 0, tberr.ErrNotFound
}

func (b *BasedSeries) Next() error {
	next := b.cur + 1
	if next >= len(b.slots) {
		return tberr.ErrIndexExceed
	}
	b.cur++
	return nil
}
func (b *BasedSeries) HasMore() bool {
	next := b.cur + 1
	return next < len(b.slots)
}

func (b *BasedSeries) Reset() {
	b.cur = 0
}
