package slotprovider

var emptyFunction = func() {}

type SlotProvider interface {
	AcquireSlot() (hasSlot bool, release func())
	OpenSlots() int
}

// returns the currently fastest SlotProvider implementation
func New(nrOfSlots int) SlotProvider {
	return NewWithMutex(nrOfSlots)
}
