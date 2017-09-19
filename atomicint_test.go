package slotprovider_test

import (
	"testing"
	"github.com/chrisport/slotprovider"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestAtUInt64_givenNoSlotOccupied_whenAcquireSlot_thenReturnTrue(t *testing.T) {
	sp := slotprovider.NewWithAtomicUInt64(nrOfSlots)

	results := make([]bool, nrOfSlots+1)
	for i := 0; i < nrOfSlots; i++ {
		results[i], _ = sp.AcquireSlot()
	}

	for i := 0; i < nrOfSlots; i++ {
		assert.True(t, results[i])
	}
}

func TestAtUInt6464_givenAllSlotOccupied_whenOneReleasedAndAcquireSlot_thenReturnTrue(t *testing.T) {
	sp := slotprovider.NewWithAtomicUInt64(nrOfSlots)
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

func TestAtUInt64_givenAllSlotsOccupied_whenAcquireSlot_thenReturnFalse(t *testing.T) {
	sp := slotprovider.NewWithAtomicUInt64(nrOfSlots)

	results := make([]bool, nrOfSlots+1)
	for i := 0; i < nrOfSlots+1; i++ {
		results[i], _ = sp.AcquireSlot()
	}
	for i := 0; i < nrOfSlots; i++ {
		assert.True(t, results[i])
	}
	assert.False(t, results[nrOfSlots])
}


func Benchmark_AtomicUInt64(b *testing.B) {
	sp := slotprovider.NewWithAtomicUInt64(nrOfSlots)

	benchmark(b, sp)
}


func BenchmarkVerify_AtomicUInt64_parallel(b *testing.B) {
	sp := slotprovider.NewWithAtomicUInt64(nrOfSlots)

	verify_parallel(b, sp)
}


func Benchmark_AtomicUInt64_parallel(b *testing.B) {
	sp := slotprovider.NewWithAtomicUInt64(nrOfSlots)

	benchmark_parallel(b, sp)
}

