package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if all digits in the integer are even.")
	var n int
	fmt.Println("Enter the number:")
	fmt.Scan(&n)
	var y int
	for z := n; z != 0; z /= 10 {
		y = z % 10
		if y%2 != 0 {
			fmt.Println("Not all digits in", n, "are even.")
			return
		}
	}
	fmt.Println("All digits in", n, "are even.")
}
