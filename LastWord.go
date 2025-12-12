package main

import (
	"fmt"
)

// Write a function that takes a string and return a string containing its last word, followed by a newline ('\n').
//
// A word is a sequence of characters delimited by spaces or by the start/end of the argument.

func LastWord(s string) string {

	result := ""
	i := len(s) - 1

	for i >= 0 && s[i] != ' ' {
		result = string(s[i]) + result
		i--
	}
	return result + "\n"

}

func main() {
	fmt.Print(LastWord("hello there"))
	fmt.Print(LastWord(""))
	fmt.Print(LastWord("hello   .........  bye"))
}

// there
//
// bye
