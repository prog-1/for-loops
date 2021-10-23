package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	var n uint
	fmt.Println("Enter the number:")
	fmt.Scan(&n)
	var y uint
	var z uint
	if n <= 0 {
		fmt.Println(n, "is not a power of 2")
	}
	for x := n; x != 1; x /= 2 {
		y = x % 2
		z++
	}
	if y == 0 {
		fmt.Println(n, "= 2^", z)
	} else {
		fmt.Println(n, "is not a power of 2")
	}
}
