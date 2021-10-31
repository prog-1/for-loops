package main

import "fmt"

func main() {
	fmt.Println("The program determines if an integer is prime.")
	fmt.Println("Enter the number:")
	var n int
	fmt.Scan(&n)
	if n<2{
		fmt.Println(n, "is not a prime number")
		return
	}
	var i int = 2
	for i = 2; i < n; i++ {
		if n%i == 0 {
			fmt.Println(n, "is not a prime number")
			return
		}
	}
	fmt.Println(n, "is a prime number")
}

// Test:
// Input: 1
// Output: 1 is not a prime number

// Test:
// Input: 0
// Output: 0 is not a prime number

// Test:
// Input: 2
// Output: 2 is a prime number

// Test:
// Input: 56
// Output: 56 is not a prime number

// Test:
// Input: 7
// Output: 7 is a prime number
