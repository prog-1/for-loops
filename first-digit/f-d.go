package main

import "fmt"

func main() {
	fmt.Println("The program prints the first digit of the given integer.")
	fmt.Println("Enter the number:")
	var x int
	fmt.Scan(&x)
	var y int
	for y = x; y/10 != 0; {
		y /= 10
	}
	fmt.Println("The first digit is", y)
}

// Test (1)
// Input: 0
// Output: The first digit is 0

// Test (2)
// Input: 256
// Output: The first digit is 2

// Test (3)
// Input: -56
// Output: The first digit is -5

// Test (4)
// Input: 11
// Output: The first digit is 1

// Test (5)
// Input: 39729
// Output: The first digit is 3

// Test (6)
// Input: -333
// Output: The first digit is -3
