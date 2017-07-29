package slotprovider

var emptyFunction = func() {}

// SlotProvider provides functionality for slot-based concurrency
type SlotProvider interface {
	// AcquireSlot is the main function to request a slot
	// hasSlot
	// 	is true if slot has been acquired
	// 	is false if no free slot was available
	// release
	// 	must be called to release acquired slot
	// 	is safe to be called multiple times and even if no slot has been acquired
	//	is not safe to be called concurrently
	AcquireSlot() (hasSlot bool, release func())
	OpenSlots() int
}

// returns the currently fastest SlotProvider implementation
func New(nrOfSlots int) SlotProvider {
	return NewWithMutex(nrOfSlots)
}
