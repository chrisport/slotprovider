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
	sp := &spChannel{openSlots: nrOfSlots, slotChan: slotChan, notifyChan: notifyChan, ctx: ctx}
	go sp.start()
	return sp
}

func (sp *spChannel) OpenSlots() int {
	return sp.openSlots
}

func (sp *spChannel) start() {
	for {
		select {
		case sp.slotChan <- true:
		case <-sp.ctx.Done():
			return
		case <-sp.notifyChan:
			sp.openSlots++
		}
	}
}

func (sp *spChannel) Release() {
	sp.notifyChan <- true
}

func (sp *spChannel) AcquireSlot() bool {
	select {
	case <-sp.slotChan:
		return true
	default:
		return false
	}
}
