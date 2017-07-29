package slotprovider

var element = struct{}{}

type spSingleChannel struct {
	slotChan chan struct{}
}

func NewWithSingleChannel(nrOfSlots int) SlotProvider {
	var slotChan = make(chan struct{}, nrOfSlots)
	for ; nrOfSlots > 0; nrOfSlots-- {
		slotChan <- element
	}
	sp := &spSingleChannel{slotChan: slotChan}
	return sp
}
func (sp *spSingleChannel) OpenSlots() int {
	return 0
}

func (sp *spSingleChannel) Release() {
	sp.slotChan <- element
}

func (sp *spSingleChannel) AcquireSlot() bool {
	select {
	case <-sp.slotChan:
		return true
	default:
		return false
	}
}
