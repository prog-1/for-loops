package main

import "fmt"

func main() {
	fmt.Println("The program determines if an integer is a perfect square.")
	fmt.Println("Enter the number:")
	var n int
	fmt.Scanln(&n)
	var k int
	for n2 := n; k*k <= n2; k++ {
		if k*k == n2 {
			fmt.Println(n, "is a perfect square.", k, "^ 2 =", n)
			return
		}
	}
	fmt.Println(n, "is not a perfect square.")
}

// Test:
//
// Input: 0
// Output: 0 is a perfect square. 0 ^ 2 = 0
//
// Input: 1
// Output: 1 is a perfect square. 1 ^ 2 = 1
//
// Input: 67
// Output: 67 is not a perfect square.
//
// Input: 144
// Output: 144 is a perfect square. 12 ^ 2 = 144
//
// Input: -16
// Output: -16 is not a perfect square.
