package main

import (
	"fmt"
)

// Write a function that takes a string and return a string containing its first word, followed by a newline ('\n').
//
// A word is a sequence of characters delimited by spaces or by the start/end of the argument.

func LastWord(s string) string {

	result := ""
	i := len(s) - 1

	for i >= 0 && s[i] == ' ' {
		i--
	}

	for i >= 0 && s[i] != ' ' {
		result = string(s[i]) + result
		i--
	}
	return result + "\n"

}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}

// hello
//
// hello
