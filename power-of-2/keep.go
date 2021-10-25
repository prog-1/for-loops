package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	fmt.Println("Enter the number:")
	var n int64
	a := 0
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("your number is not a power of 2")
	}
	for n%2 == 0 {
		n = n / 2
		a++
	}
	if n == 1 {
		fmt.Println("your number =", a, "^ 2")
	} else {
		fmt.Println("your number is not a power of 2")
	}
}
