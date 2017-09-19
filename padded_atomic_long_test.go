package slotprovider_test

import (
	"testing"
	"github.com/chrisport/slotprovider"
)

func Benchmark_AtomicPaddedLong(b *testing.B) {
	sp := slotprovider.NewWithAtomicPaddedLong(nrOfSlots)
	benchmark(b, sp)
}

func BenchmarkVerify_AtomicPaddedLong_parallel(b *testing.B) {
	sp := slotprovider.NewWithAtomicPaddedLong(nrOfSlots)
	verify_parallel(b, sp)
}

func Benchmark_AtomicPaddedLong_parallel(b *testing.B) {
	sp := slotprovider.NewWithAtomicPaddedLong(nrOfSlots)
	benchmark_parallel(b, sp)
}
