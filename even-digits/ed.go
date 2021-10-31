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

// Test
// Input: 487
// Output: Not all digits in 487 are even.

// Test
// Input: 1 
// Output: Not all digits in 487 are even.

// Test
// Input: 682
// Output: Not all digits in 487 are even. 