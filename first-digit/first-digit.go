package main

import "fmt"

func main() {
	fmt.Println("The program prints the first digit of the given integer.")
	fmt.Println("Enter the number:")
	var n int
	fmt.Scanln(&n)
	var d int
	for ; n != 0; n /= 10 {
		d = n % 10
	}
	fmt.Println("The first digit is", d)
}

// Test:
//
// Input: 1
// Output: The first digit is 1
//
// Input: 3837485
// Output : The first digit is 3
//
// Input: -15390
// Output: The first digit is -1
//
// Input: 999345721
// Output: THe first digit is 9
