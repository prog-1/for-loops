package main

import "fmt"

func main() {
	fmt.Println("The program determines if all digits in the integer are even.")
	fmt.Println("Enter the number:")
	var x int
	fmt.Scan(&x)
	for ed := x; ed != 0; ed /= 10 {
		if (ed%10)%2 != 0 {
			fmt.Println("Not digits in", x, "are even.")
			return
		}
	}
	fmt.Println("All digits in", x, "are even.")
}

// Test (1)
// Input: 0
// Output: All digits in 0 are even.

// Test (2)
// Input: 1
// Output: Not digits in 1 are even.

// Test (3)
// Input: 34
// Output: Not digits in 34 are even.

// Test (4)
// Input: 468
// Output: All digits in 468 are even.

// Test (5)
// Input: -342
// Output: Not digits in -342 are even.

// Test (6)
// Input: -682
// Output: All digits in -682 are even.
