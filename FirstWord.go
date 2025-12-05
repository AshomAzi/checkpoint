package main

import (
	"fmt"
)

// Write a function that takes a string and return a string containing its first word, followed by a newline ('\n').
//
// A word is a sequence of characters delimited by spaces or by the start/end of the argument.

func FirstWord(s string) string {

	if len(s) == 0 {
		return " "
	}

	for j, i := range s {
		if i == ' ' {
			return s[:j] + "\n"
		}
	}
	return s + "\n"

}

func main() {
	fmt.Println(FirstWord("hello there"))
	fmt.Println(FirstWord(""))
	fmt.Println(FirstWord("hello   .........  bye"))
}

// hello
//
// hello
