package main

import (
	"fmt"
)

// Write a function called RepeatAlpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.
//
// 'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...

func RepeatAlpha(s string) string {

	result := " "
	for _, i := range s {
		count := 0
		char := byte(i)

		if char >= 'a' && char <= 'z' {
			count = int(char - 'a' + 1)
		} else if char >= 'A' && char <= 'Z' {
			count = int(char - 'A' + 1)
		} else {
			count = 1
		}
		for r := 0; r < count; r++ {
			result += string(i)
		}
	}
	return result
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
