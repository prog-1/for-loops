package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program prints the first digit of the given integer.")
	var n int
	fmt.Println("Enter the number:")
	fmt.Scan(&n)
	var y int
	for z := n; z != 0; z /= 10 {
		y = z % 10
	}
	fmt.Println("The first digit is", y)
}
