package main

import "fmt"

func main() {
	fmt.Println("The program determines if an integer is prime.")
	fmt.Println("Enter the number:")
	var n int
	fmt.Scan(&n)
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
// Input: 2, 4, 199, 7, 16
// Output: 2 is a prime number, 4 is not a prime number, 199 is a prime number, 7 is a prime number, 16 is not a prime number
