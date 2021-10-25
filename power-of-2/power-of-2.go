package main

import "fmt"

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	fmt.Println("Enter the number:")
	var n int
	fmt.Scanln(&n)
	if n == 1 {
		fmt.Println(n, "= 2 ^ 0")
		return
	}
	if n <= 0 || n%2 != 0 {
		fmt.Println(n, "is not a power of 2.")
		return
	}
	var x, y int
	for n2 := n; n2 != 1; n2 /= 2 {
		x = n2 % 2
		y++
	}
	if x == 0 {
		fmt.Println(n, "= 2 ^", y)
	} else {
		fmt.Println(n, "is not a power of 2.")
	}
}

// Test:
//
// Input: 0
// Output: 0 is not a power of 2.
//
// Input: 1
// Output: 1 = 2 ^ 0
//
// Input: 4
// Output: 4 = 2 ^ 2
//
// Input: 64
// Output: 64 = 2 ^ 6
//
// Input: 15
// Output: 15 is not a power of 2.
//
// Input: -36
// Output: -36 is not a power of 2.
