package main

import (
	"strings"
	"unicode"
)

func ClearPassword(s string) string {
	result := ""

	for _, r := range s {
		if unicode.IsSpace(r) || unicode.IsControl(r) {
		} else {
			result += strings.ToUpper(string(r))
		}
	}

	return result
}
