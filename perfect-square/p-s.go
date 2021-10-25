package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is a perfect square.")
	fmt.Println("Enter the number:")
	var x int
	fmt.Scan(&x)
	for i := 0; i*i <= x; i++ {
		if i*i == x {
			fmt.Println(x, "is a perfect square.", i, "^2=", x)
			return
		}
	}
	fmt.Println(x, "is not a perfect square.")
}

// Test (1)
// Input: 0
// Output: 0 is a perfect square. 0 ^2= 0

// Test (2)
// Input: 1
// Output: 1 is a perfect square. 1 ^2= 1

// Test (3)
// Input: 9
// Output: 9 is a perfect square. 3 ^2= 9

// Test (4)
// Input: 23
// Output: 23 is not a perfect square.

// Test (5)
// Input: 3492 is not a perfect square.
// Output: 3492

// Test (1)
// Input: -121
// Output: -121 is not a perfect square.
