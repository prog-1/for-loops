package main

import "fmt"

func main() {
	fmt.Println("The program prints the first digit of the given integer.")
	fmt.Println("Enter the number:")
	var x int
	fmt.Scan(&x)
	var f int
	for f = x; f/10 != 0 ;{
		f /= 10
	}
	fmt.Println("The first digit is", f)
}

// Test:
// Input: -43
// Output: The first digit is -4

// Test:
// Input: 0
// Output: The first digit is 0

// Test:
// Input: 53822
// Output: The first digit is 5

// Test: 
// Input: 10
// Output: The first digit is 1
