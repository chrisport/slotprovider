package slotprovider

import (
	"sync/atomic"
)

type spAtomicInt struct {
	openSlots int64
}

func NewWithAtomicInt(nrOfSlots int) SlotProvider {
	sp := &spAtomicInt{openSlots: int64(nrOfSlots)}
	return sp
}

func (sp *spAtomicInt) OpenSlots() int {
	return int(atomic.LoadInt64(&sp.openSlots))
}

func (sp *spAtomicInt) release() {
	atomic.AddInt64(&sp.openSlots, 1)
}

func (sp *spAtomicInt) AcquireSlot() (hasSlot bool, release func()) {
	if atomic.AddInt64(&sp.openSlots, -1) < 0 {
		atomic.AddInt64(&sp.openSlots, 1)
		return false, emptyFunction
	}
	return true, sp.release
}
