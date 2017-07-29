package slotprovider_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/chrisport/slotprovider"
)

func TestSingleChan_givenNoSlotOccupied_whenAcquireSlot_thenReturnTrue(t *testing.T) {
	sp = slotprovider.NewWithSingleChannel(nrOfSlots)

	results := make([]bool, 11)
	for i := 0; i < 10; i++ {
		results[i], _ = sp.AcquireSlot()
	}

	for i := 0; i < 10; i++ {
		assert.True(t, results[i])
	}
}

func TestSingleChan_givenAllSlotOccupied_whenOneReleasedAndAcquireSlot_thenReturnTrue(t *testing.T) {
	sp = slotprovider.NewWithSingleChannel(nrOfSlots)
	var res bool
	var release func()
	for i := 0; i < 10; i++ {
		res, release = sp.AcquireSlot()
		assert.True(t, res)
	}

	release()
	res, _ = sp.AcquireSlot()

	assert.True(t, res)
}

func TestSingleChan_givenAllSlotsOccupied_whenAcquireSlot_thenReturnFalse(t *testing.T) {
	sp = slotprovider.NewWithSingleChannel(nrOfSlots)

	results := make([]bool, 11)
	for i := 0; i < 11; i++ {
		results[i], _ = sp.AcquireSlot()
	}

	for i := 0; i < 10; i++ {
		assert.True(t, results[i])
	}
	assert.False(t, results[10])
}
