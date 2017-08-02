package slotprovider_test

import (
	"github.com/chrisport/slotprovider"
	"testing"
	"sync"
)

var global int
var globalMux = sync.Mutex{}
var globalCounter = 0

func benchmark(b *testing.B, sp slotprovider.SlotProvider) {
	var release func()
	for i := 0; i < b.N; i++ {
		_, release = sp.AcquireSlot()
		release()
	}
	global = sp.OpenSlots()
}

func benchmark_parallel(b *testing.B, sp slotprovider.SlotProvider) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var hasSlot bool
			var release func()
			for !hasSlot {
				hasSlot, release = sp.AcquireSlot()
			}
			release()
		}
	})
	global = sp.OpenSlots()
}

// verify_parallel is to test that the # of acquired slots never exceeds the maximum number of slots even under parallel
// load. This test does not proof correctness.
func verify_parallel(b *testing.B, sp slotprovider.SlotProvider) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var hasSlot bool
			var release func()
			for !hasSlot {
				hasSlot, release = sp.AcquireSlot()
			}
			globalMux.Lock()
			globalCounter++
			if globalCounter > nrOfSlots {
				panic("More goroutines are in critical section than allowed")
			}
			globalCounter--
			globalMux.Unlock()
			release()
		}
	})
	global = sp.OpenSlots()
}
