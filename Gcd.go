package main

import (
	"fmt"
)

// Write a function that takes two uint representing two strictly positive integers and returns their greatest common divisor.
// If any of the input numbers is 0, the function should return 0.
func Gcd(a, b uint) uint {

	if a == 0 || b == 0 {
		return 0
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a

	// if a == 0 || b == 0 {
	// 	return 0
	// }

	// for b != 0 {
	// 	a, b = b, a%b
	// }
	// return a

}

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}

// 2
// 6
// 7
// 1
