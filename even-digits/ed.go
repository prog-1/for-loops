package main

import "fmt"

func main() {
	fmt.Println("The program determines if all digits in the integer are even.")
	fmt.Println("Enter the number:")
	var x int
	fmt.Scan(&x)
	for d := x; d != 0; d /= 10 {
		z := (d % 10)
		if z%2 != 0 {
			fmt.Println("Not all digits in", x, "are even.")
			return
		}
	}
	fmt.Println("All digits in", x, "are even.")
}

// Test:
// Input: 2860, 2816, 4675, 44
// Output: All digits in 2860 are even.; Not all digits in 2816 are even.; Not all digits in 4675 are even.; All digits in 44 are even.
