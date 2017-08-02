# slotprovider

Experimental package.
Manages a number of free slots which can be acquired and released concurrently.
**no waiting for free slot:** If there is no free slot, the acquire method returns immediately with false.
There are 4 implementations
 - Using int64 and atomic package
 - Using sync.Mutex
 - Using 1 buffered channel
 - Using 2 channels and go routine to manage free slots

### Benchmark results:
(MacBook Pro Early 2015, 2.7 GHz Intel Core i5)
```
Benchmark_AtomicInt-4                           10000000               122 ns/op
Benchmark_Mutex-4                               10000000               132 ns/op
Benchmark_SingleChan-4                          10000000               167 ns/op
Benchmark_MultiChan-4                            2000000               579 ns/op

Benchmark_AtomicInt_parallel-4                  20000000                88.8 ns/op
Benchmark_Mutex_parallel-4                       5000000               291 ns/op
Benchmark_SingleChan_parallel-4                  5000000               255 ns/op
Benchmark_MultiChan_parallel-4                   2000000               666 ns/op
```

## Example usage

```
nrOfSlots := 2
var sp slotprovider.SlotProvider

sp = slotprovider.NewWithAtomicInt(nrOfSlots)
// sp = slotprovider.NewWithMutex(nrOfSlots)
// sp = slotprovider.NewWithSingleChannel(nrOfSlots)
// sp = slotprovider.NewWithMultiChannel(nrOfSlots, context.Background())


hasSlot,_ := sp.AcquireSlot() // hasSlot2 = true
hasSlot2,_ := sp.AcquireSlot() // hasSlot2 = true
hasSlot3, release := sp.AcquireSlot() // hasSlot3 = false

release()

hasSlot4 := sp.AcquireSlot() // hasSlot4 = true
```
Note:
- hasSlot: indicate whether or not the requester has acquired a slot. If false, there was no free slot left.
- release: function to release the acquired slot
 - release can be called safely, even if no slot has been acquired
 - release can be called multiple times, consequent calls will be ignored silently

# Contribution

Any ideas, feedback and contributions will be appreciated.
Please implement at least the same set of tests as the other implementations.
