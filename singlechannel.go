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

func (sp *spSingleChannel) release() {
	sp.slotChan <- element
}

func (sp *spSingleChannel) AcquireSlot() (bool, func()) {
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
