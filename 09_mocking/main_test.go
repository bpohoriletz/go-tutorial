package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const write = "write"
const sleep = "sleep"

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("five", func(t *testing.T) {
		sleepTime := 5 * time.Second

		spyTime := SpyTime{}
		sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("Expected sleep for %v, but slept for %v", sleepTime, spyTime.durationSlept)
		}
	})
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)

	return
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("sllep before every ptint", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v, got %v", want, spySleepPrinter.Calls)
		}
	})
	t.Run("regular", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spySleeper := SpySleeper{}
		Countdown(&buffer, &spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("too few sleeper calls, expected %d, got %d", countDownStart, spySleeper.Calls)
		}
	})
}
