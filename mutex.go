package slotprovider

import (
	"sync"
)

func NewWithMutex(nrOfSlots int) SlotProvider {
	return &spMutex{mut: sync.Mutex{}, openSlots: nrOfSlots}
}

type spMutex struct {
	openSlots int
	mut       sync.Mutex
}

func (sp *spMutex) OpenSlots() int {
	return sp.openSlots
}

func (sp *spMutex) Release() {
	sp.mut.Lock()
	sp.openSlots++
	sp.mut.Unlock()
}

func (sp *spMutex) AcquireSlot() bool {
	res := false
	sp.mut.Lock()
	if sp.openSlots > 0 {
		sp.openSlots--
		res = true
	}
	sp.mut.Unlock()
	return res
}
