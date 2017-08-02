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
	if atomic.LoadInt64(&sp.openSlots) < 0 {
		return false, emptyFunction
	}
	newVal := atomic.AddInt64(&sp.openSlots, -1)
	if newVal < 0 {
		sp.release()
		return false, emptyFunction
	} else {
		f := sp.release
		return true, func() {
			f()
			f = emptyFunction
		}
	}
}
