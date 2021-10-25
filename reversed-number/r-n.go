package main

import "fmt"

func main() {
	fmt.Println("The program reverses the digits in the number.")
	fmt.Println("Enter the number:")
	var x uint
	fmt.Scan(&x)
	var y uint
	for rn := x; rn != 0; rn /= 10 {
		y = y*10 + rn%10
	}
	fmt.Println("The reversed number is", y)
}

// Test (1)
// Input: 0
// Output: The reversed number is 0

// Test (2)
// Input: 1
// Output: The reversed number is 1

// Test (3)
// Input: 9
// Output: The reversed number is 9

// Test (4)
// Input: 356
// Output: The reversed number is 653

// Test (5)
// Input: 47932
// Output: The reversed number is 23974
