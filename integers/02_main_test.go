package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if expected != sum {
		t.Errorf("expected %d, got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 5)
	fmt.Println(sum)
	// Output: 7
}
