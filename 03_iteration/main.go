package iteration

import "strings"

func Repeat(char string, times int) string {
	var repeated strings.Builder

	for i := 0; i < times; i++ {
		repeated.WriteString(char)
	}

	return repeated.String()
}
