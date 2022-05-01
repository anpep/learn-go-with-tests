package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter works", func(t *testing.T) {
		const times = 3
		counter := NewCounter()
		for i := 0; i < times; i++ {
			counter.Increment()
		}
		assertCounterEquals(t, counter, times)
	})

	t.Run("incrementing the counter concurrently works", func(t *testing.T) {
		const numGoroutines = 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func() {
				counter.Increment()
				wg.Done()
			}()
		}

		wg.Wait()
		assertCounterEquals(t, counter, numGoroutines)
	})
}

func assertCounterEquals(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}
