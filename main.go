package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//Задача 1

	var r rune = 'A'
	fmt.Println(r, string(r))

	so := "Go!"
	fmt.Println(len(so), utf8.RuneCountInString(so))

	ru := "Привет"
	fmt.Println(len(ru), utf8.RuneCountInString(ru))

	// Задача 2
	// var bad rune = "A"
	//.\main.go:21:17: cannot use "A" (untyped string constant) as rune value in variable declaration
	var good rune = 'A'
	fmt.Println(good, string(good))

	//Задача 3
	s := "Жук"
	fmt.Println(s[0], s[1], s[2])

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for i, r := range s {
		fmt.Println("range", i, r, string(r))
	}

	// Задача 4
	// s[0] = 'A'
	// .\main.go:39:2: cannot assign to s[0] (neither addressable nor a map index expression)

	// Задача 5
	fmt.Println(OnlyLetters("12abc"))

	// Задача 7
	fmt.Println(ReverseRunes("1dv2abc"))
	fmt.Println(ReverseRunes("Привет"))
	fmt.Println(ReverseRunes("你好"))
	fmt.Println(ReverseRunes("🙂👍"))

	// Задача 8
	fmt.Println(SubstrRunes("Интукод", 2, 4))
	fmt.Println(SubstrRunesTo("Интукод", 1, 4))

	// Задача 9. Очистка пароля
	fmt.Println(ClearPassword("  q w e\nr\t "))
	fmt.Println(ClearPassword("  пароль🙂 "))

}
