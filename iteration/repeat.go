package iteration

import "strings"

func Repeat(character string, size int) string {
	var repeated strings.Builder

	for i := 0; i < size; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}
