# slotprovider

## SlotProvider
Manages a number of free slots

## Usage
```
nrOfSlots := 2
sp = slotprovider.NewSlotProvider(5, context.Background())

hasSlot, release := sp.AcquireSlot() // hasSlot2 = true
hasSlot2, _ := sp.AcquireSlot() // hasSlot2 = true
hasSlot3, _ := sp.AcquireSlot() // hasSlot3 = false

release()

hasSlot4, _ := sp.AcquireSlot() // hasSlot4 = true

```
