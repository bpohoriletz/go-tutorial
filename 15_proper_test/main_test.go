package propertest

import (
	"strconv"
	"testing"
	"testing/quick"
)

func TestPropertyConsecutive(t *testing.T) {
	assertion := func(num uint16) bool {
		if num > 3999 {
			return true
		}
		roman := ConvertToRoman(num)
		t.Log(roman)

		maxConsecutive := 0
		consec := 0

		charWas := ""

		for i := range roman {
			if string(roman[i]) == charWas {
				consec += 1
			} else {
				maxConsecutive = max(consec, maxConsecutive)
				consec = 0
			}
			charWas = string(roman[i])
		}

		return maxConsecutive < 4
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed consecutive checks", err)
	}
}

func TestPropertyComposition(t *testing.T) {
	assertion := func(num uint16) bool {
		roman := ConvertToRoman(num)
		arab := ConvertToArabic(roman)

		return num == arab
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed composition checks", err)
	}
}

func TestRomanConvertion(t *testing.T) {
	tests := []struct {
		arabic uint16
		roman  string
	}{
		{1, "I"}, {2, "II"}, {3, "III"}, {4, "IV"}, {5, "V"}, {6, "VI"}, {7, "VII"}, {8, "VIII"}, {9, "IX"}, {10, "X"},
		{11, "XI"}, {12, "XII"}, {13, "XIII"}, {14, "XIV"}, {15, "XV"}, {16, "XVI"}, {17, "XVII"}, {18, "XVIII"}, {19, "XIX"}, {20, "XX"},
		{41, "XLI"}, {42, "XLII"}, {43, "XLIII"}, {44, "XLIV"}, {45, "XLV"}, {46, "XLVI"}, {47, "XLVII"}, {48, "XLVIII"}, {49, "XLIX"}, {50, "L"},
		{510, "DX"}, {463, "CDLXIII"}, {275, "CCLXXV"},
		{1984, "MCMLXXXIV"}, {2025, "MMXXV"},
	}

	for _, test := range tests {
		t.Run(test.roman, func(t *testing.T) {
			got := ConvertToRoman(test.arabic)
			assertEqual(t, got, test.roman)
		})
	}

	for _, test := range tests {
		t.Run(test.roman, func(t *testing.T) {
			got := ConvertToArabic(test.roman)
			assertEqual(t, strconv.FormatUint(uint64(got), 10), strconv.FormatUint(uint64(test.arabic), 10))
		})
	}
}

func assertEqual(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
