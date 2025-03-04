package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"time"
)

const countDownStart = 3
const finalWord = "Go!"

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type Sleeper interface {
	Sleep()
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := range countDownFrom(3) {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	sleeper := ConfigurableSleeper{1 * time.Second, time.Sleep}

	Countdown(os.Stdout, sleeper)
}
