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
	for d := b; d != 0; d /= 10 {
		c = d % 10
	}
	fmt.Println("The first digit is", d)
}