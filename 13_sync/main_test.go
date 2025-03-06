package ssync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("count 3", func(t *testing.T) {
		counter := Counter{}
		for range 3 {
			counter.Inc()
		}

		assertCounter(t, &counter, 3)
	})
	t.Run("concurrent 1000", func(t *testing.T) {
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(1000)

		for range 1000 {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, &counter, 1000)
	})
}

func assertCounter(t *testing.T, counter *Counter, expect int) {
	if expect == counter.Value() {
		t.Errorf("got %q, want %d", counter, expect)
	}
}
