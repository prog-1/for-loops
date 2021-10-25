package main

import "fmt"

func main() {
	fmt.Println("The program determines if all digits in the integer are even.")
	fmt.Println("Enter the number:")
	var n int
	fmt.Scanln(&n)
	var x int
	for n2 := n; n2 != 0; n2 /= 10 {
		x = n2 % 10
		if x%2 != 0 {
			fmt.Println("Not all digits in", n, "are even.")
			return
		}
	}
	fmt.Println("All digits in", n, "are even.")
}

// Test:
//
// Input: 0
// Output: All digits in 0 are even.
//
// Input: 1
// Output: Not all digits in 1 are even.
//
// Input: 48026
// Output: All digits in 48026 are even.
//
// Input: 389412
// Output: Not all digits in 389412 are even.
//
// Input: -4802
// Output: All digits in -4802 are even.
