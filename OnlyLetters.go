package main

import "unicode"

func OnlyLetters(s string) string {
	result := ""
	for _, r := range s {
		if unicode.IsLetter(r) {
			result += string(r)
		}
	}

	return result
}
