package main

import "fmt"

func SubstrRunes(s string, start, length int) string {
	ru := []rune(s)

	if len(ru)-1 < start || len(ru)-1 < length {
		return "error"
	}

	result := ""

	for i := start; i < start+length; i++ {
		fmt.Println(i)
		result += string(ru[i])
	}

	return result
}

func SubstrRunesTo(s string, start, length int) string {
	ru := []rune(s)

	if len(ru)-1 < start || len(ru)-1 < length {
		return "error"
	}

	return string((ru[start : start+length]))
}
