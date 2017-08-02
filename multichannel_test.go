package slotprovider_test

import (
	"golang.org/x/net/context"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/chrisport/slotprovider"
	"time"
)

const nrOfSlots = 137

var sp slotprovider.SlotProvider

func setupProvider(nrOfSlots int) (cancelFunc func()) {
	ctx, cancel := context.WithCancel(context.Background())
	sp = slotprovider.NewWithMultiChannel(nrOfSlots, ctx)
	return cancel
}

func Test_givenNoSlotOccupied_whenAcquireSlot_thenReturnTrue(t *testing.T) {
	defer setupProvider(nrOfSlots)()

	results := make([]bool, nrOfSlots+1)
	for i := 0; i < nrOfSlots; i++ {
		results[i], _ = sp.AcquireSlot()
	}

	for i := 0; i < nrOfSlots; i++ {
		assert.True(t, results[i])
	}
}

func Test_givenAllSlotOccupied_whenOneReleasedAndAcquireSlot_thenReturnTrue(t *testing.T) {
	defer setupProvider(nrOfSlots)()
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

func Test_givenAllSlotsOccupied_whenAcquireSlot_thenReturnFalse(t *testing.T) {
	defer setupProvider(nrOfSlots)()

	results := make([]bool, nrOfSlots+1)
	for i := 0; i < nrOfSlots+1; i++ {
		results[i], _ = sp.AcquireSlot()
	}
	for i := 0; i < nrOfSlots; i++ {
		assert.True(t, results[i])
	}
	assert.False(t, results[nrOfSlots])
}

func Benchmark_MultiChan(b *testing.B) {
	defer setupProvider(nrOfSlots)()

	benchmark(b, sp)
}

func Benchmark_MultiChan_parallel(b *testing.B) {
	defer setupProvider(nrOfSlots)()

	benchmark_parallel(b, sp)

}

func BenchmarkVerify_MultiChan_parallel(b *testing.B) {
	defer setupProvider(nrOfSlots)()
	verify_parallel(b, sp)
}