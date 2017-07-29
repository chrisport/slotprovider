package slotprovider_test

import (
	"golang.org/x/net/context"
	"testing"
	"github.com/stretchr/testify/assert"
	"com/github/chrisport/slotprovider"
)

const nrOfSlots = 10
var sp slotprovider.SlotProvider

func setupProvider() (cancelFunc func()) {
	ctx, cancel := context.WithCancel(context.Background())
	sp = slotprovider.NewSlotProvider(nrOfSlots,ctx)
	return cancel
}

func Test_givenNoSlotOccupied_whenAcquireSlot_thenReturnTrue(t *testing.T) {
	defer setupProvider()()

	results := make([]bool, 11)
	for i := 0; i < 10; i++ {
		results[i], _ = sp.AcquireSlot()
	}

	for i := 0; i < 10; i++ {
		assert.True(t, results[i])
	}
}

func Test_givenAllSlotOccupied_whenOneReleasedAndAcquireSlot_thenReturnTrue(t *testing.T) {
	defer setupProvider()()
	var release func()
	var res bool
	for i := 0; i < 10; i++ {
		res, release = sp.AcquireSlot()
		assert.True(t,res)
	}

	release()
	res, _ = sp.AcquireSlot()

	assert.True(t, res)
}

func Test_givenAllSlotsOccupied_whenAcquireSlot_thenReturnFalse(t *testing.T) {
	defer setupProvider()()

	results := make([]bool, 11)
	for i := 0; i < 11; i++ {
		results[i], _ = sp.AcquireSlot()
	}

	for i := 0; i < 10; i++ {
		assert.True(t, results[i])
	}
	assert.False(t, results[10])
}
