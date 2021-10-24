package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program prints the first digit of the given integer.")
	fmt.Println("Enter the number:")
	var b int
	fmt.Scanln(&b)
	var c int
	for c = b; c/10 != 0 {
		c /= 10
	}
	fmt.Println("First digit is", c)
}