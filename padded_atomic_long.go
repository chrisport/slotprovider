package slotprovider

import "sync/atomic"

type spPaddedAtomicLong struct {
	openSlots              int64
	p1, p2, p3, p4, p5, p6 int64
}

func NewWithAtomicPaddedLong(nrOfSlots int) SlotProvider {
	sp := &spPaddedAtomicLong{
		openSlots: int64(nrOfSlots),
		p1:        7,
		p2:        7,
		p3:        7,
		p4:        7,
		p5:        7,
		p6:        7,
	}
	return sp
}

func (sp *spPaddedAtomicLong) OpenSlots() int {
	return int(atomic.LoadInt64(&sp.openSlots))
}

func (sp *spPaddedAtomicLong) release() {
	atomic.AddInt64(&sp.openSlots, 1)
}

func (sp *spPaddedAtomicLong) AcquireSlot() (hasSlot bool, release func()) {
	if atomic.AddInt64(&sp.openSlots, -1) <= 0 {
		atomic.AddInt64(&sp.openSlots, 1)
		return false, emptyFunction
	}
	return true, sp.release
}
