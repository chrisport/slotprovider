package slotprovider

import (
	"sync/atomic"
)
const minusOne uint64 = ^uint64(0)

type spAtomicUInt64 struct {
	openSlots uint64
}

func NewWithAtomicUInt64(nrOfSlots int) SlotProvider {
	sp := &spAtomicUInt64{
		openSlots: uint64(nrOfSlots),
	}
	return sp
}

func (sp *spAtomicUInt64) OpenSlots() int {
	return int(atomic.LoadUint64(&sp.openSlots))
}

func (sp *spAtomicUInt64) release() {
	atomic.AddUint64(&sp.openSlots, 1)
}

func (sp *spAtomicUInt64) AcquireSlot() (hasSlot bool, release func()) {
	//precheck to prevent deadlock
	if atomic.LoadUint64(&sp.openSlots) <= 0 {
		return false, emptyFunction
	}
	newVal := atomic.AddUint64(&sp.openSlots, minusOne)
	if newVal < 0 {
		sp.release()
		return false, emptyFunction
	} else {
		return true, sp.release
	}
}
