package slotprovider

import (
	"golang.org/x/net/context"
)

type spChannel struct {
	openSlots  int
	slotChan   chan bool
	notifyChan chan bool
	ctx        context.Context
}

func NewWithMultiChannel(nrOfSlots int, ctx context.Context) SlotProvider {
	var slotChan = make(chan bool, nrOfSlots)
	var notifyChan = make(chan bool)
	openSlots := nrOfSlots
	for ; openSlots > 0; openSlots-- {
		slotChan <- true
	}
	sp := &spChannel{openSlots: 0, slotChan: slotChan, notifyChan: notifyChan, ctx: ctx}
	go sp.start()
	return sp
}

func (sp *spChannel) OpenSlots() int {
	return sp.openSlots
}

func (sp *spChannel) start() {
	for {
		for ; sp.openSlots > 0; sp.openSlots-- {
			sp.slotChan <- true
		}

		<-sp.notifyChan
		sp.openSlots++
	}
}

func (sp *spChannel) release() {
	sp.notifyChan <- true
}

func (sp *spChannel) AcquireSlot() (hasSlot bool, release func()) {
	select {
	case <-sp.slotChan:
		f := sp.release
		return true, func() {
			f()
			f = emptyFunction
		}
	default:
		return false, emptyFunction
	}
}
