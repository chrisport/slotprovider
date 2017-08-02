package slotprovider_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/chrisport/slotprovider"
	"time"
)

func TestAtInt_givenNoSlotOccupied_whenAcquireSlot_thenReturnTrue(t *testing.T) {
	sp := slotprovider.NewWithAtomicInt(nrOfSlots)

	results := make([]bool, nrOfSlots+1)
	for i := 0; i < nrOfSlots; i++ {
		results[i], _ = sp.AcquireSlot()
	}

	for i := 0; i < nrOfSlots; i++ {
		assert.True(t, results[i])
	}
}

func TestAtInt_givenAllSlotOccupied_whenOneReleasedAndAcquireSlot_thenReturnTrue(t *testing.T) {
	sp := slotprovider.NewWithAtomicInt(nrOfSlots)
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

func TestAtInt_givenAllSlotsOccupied_whenAcquireSlot_thenReturnFalse(t *testing.T) {
	sp := slotprovider.NewWithAtomicInt(nrOfSlots)

	results := make([]bool, nrOfSlots+1)
	for i := 0; i < nrOfSlots+1; i++ {
		results[i], _ = sp.AcquireSlot()
	}
	for i := 0; i < nrOfSlots; i++ {
		assert.True(t, results[i])
	}
	assert.False(t, results[nrOfSlots])
}


func Benchmark_AtomicInt(b *testing.B) {
	sp := slotprovider.NewWithAtomicInt(nrOfSlots)

	benchmark(b, sp)
}


func BenchmarkVerify_AtomicInt_parallel(b *testing.B) {
	sp := slotprovider.NewWithAtomicInt(nrOfSlots)

	verify_parallel(b, sp)
}


func Benchmark_AtomicInt_parallel(b *testing.B) {
	sp := slotprovider.NewWithAtomicInt(nrOfSlots)

	benchmark_parallel(b, sp)
}

