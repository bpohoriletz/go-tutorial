package structs_methods_interfaces

import (
	"math"
	"testing"
)

func TestAreaTable(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{Height: 1, Width: 2}, 2},
		{"Circle", Circle{Radius: 10}, 10 * 10 * math.Pi},
		{"Triangle", Triangle{Base: 3, Height: 4}, (3 * 4) / 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()

			if got != test.want {
				t.Errorf("%#v want %g, got %g", test.shape, test.want, got)
			}

		})
	}
}

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{1.21, 2.14})
	want := 2 * (1.21 + 2.14)

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("want %.2f, got %.2f", want, got)
		}

	}
	t.Run("Rectangle", func(t *testing.T) {
		checkArea(t, Rectangle{2.1, 3.3}, 2.1*3.3)
	})

	t.Run("Circle", func(t *testing.T) {
		checkArea(t, Circle{12.1}, math.Pi*12.1*12.1)
	})
}
