package slotprovider_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/chrisport/slotprovider"
)

func TestSingleChan_givenNoSlotOccupied_whenAcquireSlot_thenReturnTrue(t *testing.T) {
	sp = slotprovider.NewWithSingleChannel(nrOfSlots)

	results := make([]bool, nrOfSlots+1)
	for i := 0; i < nrOfSlots; i++ {
		results[i], _ = sp.AcquireSlot()
	}

	for i := 0; i < nrOfSlots; i++ {
		assert.True(t, results[i])
	}
}

func TestSingleChan_givenAllSlotOccupied_whenOneReleasedAndAcquireSlot_thenReturnTrue(t *testing.T) {
	sp = slotprovider.NewWithSingleChannel(nrOfSlots)
	var res bool
	var release func()
	for i := 0; i < nrOfSlots; i++ {
		res, release = sp.AcquireSlot()
		assert.True(t, res)
	}

	release()
	res, _ = sp.AcquireSlot()

	assert.True(t, res)
}

func TestSingleChan_givenAllSlotsOccupied_whenAcquireSlot_thenReturnFalse(t *testing.T) {
	sp = slotprovider.NewWithSingleChannel(nrOfSlots)

	results := make([]bool, nrOfSlots+1)
	for i := 0; i < nrOfSlots+1; i++ {
		results[i], _ = sp.AcquireSlot()
	}

	for i := 0; i < nrOfSlots; i++ {
		assert.True(t, results[i])
	}
	assert.False(t, results[nrOfSlots])
}

func Benchmark_SingleChan(b *testing.B) {
	sp := slotprovider.NewWithSingleChannel(nrOfSlots)
	benchmark(b, sp)
}


func Benchmark_SingleChan_parallel(b *testing.B) {
	sp := slotprovider.NewWithSingleChannel(nrOfSlots)

	benchmark_parallel(b, sp)
}

func BenchmarkVerify_SingleChan_parallel(b *testing.B) {
	sp := slotprovider.NewWithSingleChannel(nrOfSlots)
	verify_parallel(b, sp)
}