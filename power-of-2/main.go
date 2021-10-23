package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	var n int
	fmt.Println("Enter the number:")
	fmt.Scan(&n)
	var y int
	var z int
	if n == 1 {
		fmt.Println(n, "=2^0")
		return
	}
	if n <= 0 || n%2 != 0 {
		fmt.Println(n, "is not a power of 2")
		return
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
