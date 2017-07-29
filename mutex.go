package slotprovider

import (
	"sync"
)

func NewWithMutex(nrOfSlots int) *SpMutex {
	return &SpMutex{Mut: sync.Mutex{}, openSlots: nrOfSlots}
}

type SpMutex struct {
	openSlots int
	Mut       sync.Mutex
}

func (sp *SpMutex) OpenSlots() int {
	return sp.openSlots
}

func (sp *SpMutex) release() {
	sp.Mut.Lock()
	sp.openSlots++
	sp.Mut.Unlock()
}

func (sp *SpMutex) AcquireSlot() (bool, func()) {
	res := false
	fun := emptyFunction
	sp.Mut.Lock()
	if sp.openSlots > 0 {
		sp.openSlots--
		res = true
		f := sp.release
		fun = func() {
			f()
			f = emptyFunction
		}
	}
	sp.Mut.Unlock()
	return res, fun
}
