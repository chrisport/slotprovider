package slotprovider

import "sync/atomic"

type spPaddedAtomicLong struct {
	padding   [8]uint64
	openSlots uint64
}

func NewWithAtomicUInt64Padded(nrOfSlots uint64) SlotProvider {
	sp := &spPaddedAtomicLong{
		openSlots: uint64(nrOfSlots),
		padding:  [8]uint64{},
	}
	return sp
}

func (sp *spPaddedAtomicLong) OpenSlots() int {
	return int(atomic.LoadUint64(&sp.openSlots))
}

func (sp *spPaddedAtomicLong) release() {
	atomic.AddUint64(&sp.openSlots, 1)
}

func (sp *spPaddedAtomicLong) AcquireSlot() (hasSlot bool, release func()) {
	if atomic.LoadUint64(&sp.openSlots) <= 0 {
		return false, emptyFunction
	}
	if atomic.AddUint64(&sp.openSlots, minusOne) < 0 {
		atomic.AddUint64(&sp.openSlots, 1)
		return false, emptyFunction
	}
	return true, sp.release
}
