package slotprovider_test

import (
	"testing"
	"github.com/chrisport/slotprovider"
	"github.com/stretchr/testify/assert"
	"time"
)


func TestAtPaddedInt64_givenNoSlotOccupied_whenAcquireSlot_thenReturnTrue(t *testing.T) {
	sp := slotprovider.NewWithAtomicPaddedLong(nrOfSlots)

	results := make([]bool, nrOfSlots+1)
	for i := 0; i < nrOfSlots; i++ {
		results[i], _ = sp.AcquireSlot()
	}

	for i := 0; i < nrOfSlots; i++ {
		assert.True(t, results[i])
	}
}

func TestAtPaddedInt64_givenAllSlotOccupied_whenOneReleasedAndAcquireSlot_thenReturnTrue(t *testing.T) {
	sp := slotprovider.NewWithAtomicPaddedLong(nrOfSlots)
	var res bool
	var release func()
	for i := 0; i < nrOfSlots; i++ {
		res, release = sp.AcquireSlot()
		assert.True(t, res)
	}

	release()
	time.Sleep(time.Second)
	res, _ = sp.AcquireSlot()

	assert.True(t, res)
}

func TestAtPaddedInt64_givenAllSlotsOccupied_whenAcquireSlot_thenReturnFalse(t *testing.T) {
	sp := slotprovider.NewWithAtomicPaddedLong(nrOfSlots)

	results := make([]bool, nrOfSlots+1)
	for i := 0; i < nrOfSlots+1; i++ {
		results[i], _ = sp.AcquireSlot()
	}
	for i := 0; i < nrOfSlots; i++ {
		assert.True(t, results[i])
	}
	assert.False(t, results[nrOfSlots])
}


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
