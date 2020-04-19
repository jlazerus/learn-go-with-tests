package iteration

import (
	"strings"
	"unicode"
)

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func MyTrimFunc(s string) string {
	t := strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsPunct(r)
	})
	return t
}
