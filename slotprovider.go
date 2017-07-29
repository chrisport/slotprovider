package slotprovider

import (
	"golang.org/x/net/context"
	"log"
	"sync"
)

type SlotProvider interface {
	AcquireSlot() bool
	Release()
	OpenSlots() int
}

type slotProvider struct {
	openSlots  int
	slotChan   chan bool
	notifyChan chan bool
	ctx        context.Context
}

func NewWithMutex(nrOfSlots int) SlotProvider {
	return &slotProviderMut{mut: sync.Mutex{}, openSlots: nrOfSlots}
}

func NewWithChannel(nrOfSlots int, ctx context.Context) SlotProvider {
	var slotChan = make(chan bool, nrOfSlots)
	var notifyChan = make(chan bool)
	openSlots := nrOfSlots
	for ; openSlots > 0; openSlots-- {
		slotChan <- true
	}
	sp := &slotProvider{openSlots: nrOfSlots, slotChan: slotChan, notifyChan: notifyChan, ctx: ctx}
	go sp.start()
	return sp
}

func (sp *slotProvider) OpenSlots() int {
	return sp.openSlots
}

func (sp *slotProvider) start() {
	for {
		select {
		case sp.slotChan <- true:
		case <-sp.ctx.Done():
			log.Println("SlotProvider shutdown")
			return
		case <-sp.notifyChan:
			sp.openSlots++
		}
	}
}

func (sp *slotProvider) Release() {
	sp.notifyChan <- true
}

func (sp *slotProvider) AcquireSlot() bool {
	select {
	case <-sp.slotChan:
		return true
	default:
		return false
	}
}

type slotProviderMut struct {
	openSlots int
	mut       sync.Mutex
}

func (sp *slotProviderMut) OpenSlots() int {
	return sp.openSlots
}

func (sp *slotProviderMut) Release() {
	sp.mut.Lock()
	sp.openSlots++
	sp.mut.Unlock()
}

func (sp *slotProviderMut) AcquireSlot() bool {
	res := false
	sp.mut.Lock()
	if sp.openSlots > 0 {
		sp.openSlots--
		res = true
	}
	sp.mut.Unlock()
	return res
}
