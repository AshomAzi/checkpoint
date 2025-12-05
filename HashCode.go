package main

import (
	"fmt"
)

// Write a function called HashCode() that takes a string as an argument and returns a new hashed string.
//
// The hash equation is computed as follows:
//
// (ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.
//
// If the resulting character is unprintable add 33 to it.

func HashCode(dec string) string {

	hash := ""

	for _, i := range dec {
		if s := int(i) + len(dec)%127; s < 33 {
			hash = hash + string(s+33)
		} else {
			hash = hash + string(s)
		}
	}
	return hash
}
func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

// B
// CD
// EDF
// Spwwz+bz}wo
