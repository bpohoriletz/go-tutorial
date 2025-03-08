package propertest

import (
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToArabic(val string) (total uint16) {
	for _, num := range allRomanNumerals {
		for strings.HasPrefix(val, num.Symbol) {
			total += num.Value
			val = strings.TrimPrefix(val, num.Symbol)
		}
	}

	return total
}

func ConvertToRoman(val uint16) string {
	var result strings.Builder

	for _, num := range allRomanNumerals {
		for val >= num.Value {
			result.WriteString(num.Symbol)
			val -= num.Value
		}
	}

	return result.String()
}
