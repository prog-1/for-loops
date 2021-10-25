package main

import "fmt"

func main() {
	fmt.Println("The program determines if an integer is prime.")
	fmt.Println("Enter the number:")
	var n int
	fmt.Scanln(&n)
	if n <= 1 {
		fmt.Println(n, "is not a prime number")
		return
	}
	for x := 2; x < n; x++ {
		if n%x == 0 {
			fmt.Println(n, "is not a prime number.")
			return
		}
	}
	fmt.Println(n, "is a prime number.")
}

// Test:
//
// Input: 0
// Output: 0 is not a prime number.
//
// Input: 1
// Output: 1 is not a prime number.
//
// Input: 5
// Output: 5 is a prime number.
//
// Input: 6
// Output: 6 is not a prime number.
//
// Input: 97
// Output: 97 is a prime number.
//
// Input: -2
// Output: -2 is not a prime number.
