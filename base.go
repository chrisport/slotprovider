package slotprovider

type SlotProvider interface {
	AcquireSlot() bool
	Release()
	OpenSlots() int
}

// returns the currently fastest SlotProvider implementation
func New(nrOfSlots int) SlotProvider {
	return NewWithMutex(nrOfSlots)
}
