package main

func ReverseRunes(s string) string {
	r := []rune(s)
	result := ""
	for i := len(r) - 1; i >= 0; i-- {
		result += string(r[i])
	}

	return result
}
