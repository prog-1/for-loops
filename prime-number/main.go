package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is prime.")
	var n int
	fmt.Println("Enter the number:")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println(n, "is not a prime number")
		return
	}
	for a := 2; a < n; a++ {
		if n%a == 0 {
			fmt.Println(n, "is not a prime number")
			return
		}
	}
	fmt.Println(n, "is a prime number")
}
