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
	for n%2 == 0 && n > 2 || n%3 == 0 && n > 3 || n%5 == 0 && n > 5 || n%7 == 0 && n > 7 {
		fmt.Println(n, "is not a prime number")
		fmt.Println("Enter the number:")
		fmt.Scan(&n)
	}
	if n <= 0 {
		fmt.Println(n, "is not a prime number")
		return
	}
	fmt.Println(n, "is a prime number")
}
