package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondHandPoint(t *testing.T) {
	tests := []struct {
		name  string
		at    time.Time
		point Point
	}{
		{"0 s", sampleTime(0, 0, 0), Point{0, 1}},
		{"30 s", sampleTime(0, 0, 30), Point{0, -1}},
		{"45 s", sampleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := secondHandPoint(test.at)

			if !roughlyEqualPoints(got, test.point) {
				t.Errorf("Expect %v, got %v", test.point, got)
			}
		})
	}
}

func TestSecondsInRadians(t *testing.T) {
	tests := []struct {
		name string
		at   time.Time
		rad  float64
	}{
		{name: "00 sec", at: sampleTime(0, 0, 0), rad: 0},
		{name: "30 sec", at: sampleTime(0, 0, 30), rad: math.Pi},
		{name: "45 sec", at: sampleTime(0, 0, 45), rad: math.Pi / 30 * 45},
		{name: "07 sec", at: sampleTime(0, 0, 7), rad: math.Pi / 30 * 7},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := secondsInRadians(test.at)

			if test.rad != got {
				t.Errorf("Got %v, want %v", got, test.rad)
			}
		})
	}
}

func roughlyEqualPoints(a, b Point) bool {
	return roughtlyEqualFloat64(a.X, b.X) &&
		roughtlyEqualFloat64(a.Y, b.Y)
}

func roughtlyEqualFloat64(a, b float64) bool {
	const equalityTreshhold = 1e-7

	return math.Abs(a-b) < equalityTreshhold
}

func sampleTime(h, m, s int) time.Time {
	return time.Date(1337, time.January, 1, h, m, s, 0, time.UTC)
}
