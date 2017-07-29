# slotprovider

## SlotProvider
Experimental package.
Manages a number of free slots which can be acquired and released concurrently.
There are 2 implementations, one using Mutex and one using channels.
Mutex implementation is 30x faster than channel.
Any feedback or ideas are welcome. (issue or PM)

## Usage
Usage of channel based implementation
```
nrOfSlots := 2
sp = slotprovider.NewSlotProvider(nrOfSlots, context.Background())
// alternative mutex based implementation:
sp = slotprovider.NewSlotProviderMut(nrOfSlots)

hasSlot := sp.AcquireSlot() // hasSlot2 = true
hasSlot2 := sp.AcquireSlot() // hasSlot2 = true
hasSlot3 := sp.AcquireSlot() // hasSlot3 = false

sp.release()

hasSlot4 := sp.AcquireSlot() // hasSlot4 = true
```