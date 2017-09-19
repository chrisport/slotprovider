# slotprovider

Experimental package.
Manages a number of free slots which can be acquired and released concurrently.   
**No blocking:** If there is no free slot, the Acquire method returns immediately with false.   
There are 5 implementations .
 - Using int64 with padding and atomic package
 - Using int64 and atomic package
 - Using sync.Mutex
 - Using 1 buffered channel
 - Using 2 channels and go routine to manage free slots

### Benchmark results:
(MacBook Pro Early 2015, 2.7 GHz Intel Core i5)
```
godep go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/chrisport/slotprovider
Benchmark_AtomicUInt64-4                                30000000                50.4 ns/op
Benchmark_AtomicUInt64Padded-4                          20000000                52.6 ns/op
Benchmark_Mutex-4                                       20000000                74.2 ns/op
Benchmark_SingleChan-4                                  10000000               100 ns/op

Benchmark_AtomicUInt64_parallel-4                       30000000                54.2 ns/op
Benchmark_AtomicUInt64Padded_parallel-4                 30000000                57.2 ns/op
Benchmark_Mutex_parallel-4                              10000000               188 ns/op
Benchmark_SingleChan_parallel-4                         10000000               210 ns/op
```

## Example usage

```
nrOfSlots := 2
var sp slotprovider.SlotProvider

sp = slotprovider.NewWithAtomicUInt64(nrOfSlots)
// sp = slotprovider.NewWithAtomicUInt64Padded(nrOfSlots)
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
  - a release function can be called safely, even if no slot has been acquired
  - a release function must not be called multiple times, this might create new slots.

# Contribution

Any ideas, feedback and contributions will be appreciated.
Please implement at least the same set of tests as the other implementations.
