package slotprovider

import (
	"time"
	"golang.org/x/net/context"
	"log"
	"sync"
)

type SlotProvider struct {
	openSlots  int
	slotChan   chan bool
	notifyChan chan bool
	ctx        context.Context
}

func NewSlotProvider(nrOfSlots int, ctx context.Context) SlotProvider {
	var slotChan = make(chan bool, nrOfSlots)
	var notifyChan = make(chan bool)
	openSlots := nrOfSlots
	for ; openSlots > 0; openSlots-- {
		slotChan <- true
	}
	return SlotProvider{openSlots: nrOfSlots, slotChan: slotChan, notifyChan: notifyChan, ctx: ctx}
}

func (sp *SlotProvider) Start() {
	for {
		for ; sp.openSlots > 0; sp.openSlots-- {
			sp.slotChan <- true
		}
		select {
		case <-sp.ctx.Done():
			log.Println("SlotProvider shutdown")
			return
		case <-sp.notifyChan:
			sp.openSlots++
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func (sp *SlotProvider) RequestSlot() (bool, func()) {
	select {
	case <-sp.slotChan:
		once := sync.Once{}
		return true, func() {
			once.Do(func() { sp.notifyChan <- true })
		}
	default:
		return false, nil
	}
}
