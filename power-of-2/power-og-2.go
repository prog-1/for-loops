package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	fmt.Print("Enter the number:")
	var a int
	fmt.Scan(&a)
	var i int
	result := 2
	i = 1
	for {
		if a == result {
			fmt.Print("2^", i, " == ", a)
			return
		}
		if a < result {
			break
		}
		i += 1
		result *= 2

	}
	fmt.Print(a, " is not a power of 2.")
}
