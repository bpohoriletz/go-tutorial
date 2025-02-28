package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat five times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})
	t.Run("Repeat one time", func(t *testing.T) {
		got := Repeat("a", 1)
		want := "a"

		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})
}

func BenchmarkRepeatFive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", 9)
	}
}

func ExampleRepeat() {
	str := Repeat("b", 3)
	fmt.Println(str)
	// Output: bbb
}
