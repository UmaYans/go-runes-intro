package main

import (
	"strings"
	"unicode"
)

func ClearPassword(s string) string {
	runes := []rune(s)
	result := ""

	for _, r := range runes {
		if unicode.IsSpace(r) || unicode.IsControl(r) {
		} else {
			result += strings.ToUpper(string(r))
		}
	}

	return result
}
