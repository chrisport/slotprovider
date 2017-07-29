package slotprovider

import (
	"sync"
)

func NewWithMutex(nrOfSlots int) *spMutex {
	return &spMutex{mut: sync.Mutex{}, openSlots: nrOfSlots}
}

type spMutex struct {
	openSlots int
	mut       sync.Mutex
}

func (sp *spMutex) OpenSlots() int {
	return sp.openSlots
}

func (sp *spMutex) release() {
	sp.mut.Lock()
	sp.openSlots++
	sp.mut.Unlock()
}

func (sp *spMutex) AcquireSlot() (bool, func()) {
	res := false
	fun := emptyFunction
	sp.mut.Lock()
	if sp.openSlots > 0 {
		sp.openSlots--
		res = true
		f := sp.release
		fun = func() {
			f()
			f = emptyFunction
		}
	}
	sp.mut.Unlock()
	return res, fun
}
