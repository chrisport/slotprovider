package slotprovider_test

import (
	"testing"
	"github.com/chrisport/slotprovider"
)

var Global int

func Benchmark_Mutex(b *testing.B) {
	sp := slotprovider.NewWithMutex(nrOfSlots)
	for i := 0; i < b.N; i++ {
		hasSlot := sp.AcquireSlot()
		if hasSlot != true {
			panic("was not true")
		}
		sp.Release()
	}
	Global = sp.OpenSlots()
}

func Benchmark_SingleChan(b *testing.B) {
	sp := slotprovider.NewWithSingleChannel(nrOfSlots)
	for i := 0; i < b.N; i++ {
		hasSlot := sp.AcquireSlot()
		if hasSlot != true {
			panic("was not true")
		}
		sp.Release()
	}
	Global = sp.OpenSlots()
}

func Benchmark_MultiChan(b *testing.B) {
	defer setupProvider()()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasSlot := sp.AcquireSlot()
		if hasSlot != true {
			panic("was not true")
		}
		sp.Release()
	}
	Global = sp.OpenSlots()
}
