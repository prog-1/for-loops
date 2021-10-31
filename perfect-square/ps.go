package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is a perfect square.")
	fmt.Println("Enter the number:")
	var x int
	fmt.Scan(&x)
	if x < 0 {
		fmt.Println("The negative number can't be perfect square")
	}
	for y := 1; y*y <= x; y++ {
		if y*y == x {
			fmt.Println(x, "is a perfect square.", y, "^ 2 =", x)
			return
		}

	}
	fmt.Println(x, "is not a perfect square.")
}

// Test:
// Input: 9, 8, 45, 121
// Output: 9 is a perfect square. 3 ^ 2 = 9, 8 is not a perfect square., 45 is not a perfect square., 121 is a perfect square. 11 ^ 2 = 121
