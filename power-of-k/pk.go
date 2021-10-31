package main

import "fmt"

func main() {
	fmt.Println("The program determines if a number n is a power of k.")
	fmt.Println("Enter n and k:")
	var n, k int
	fmt.Scan(&n, &k)
	var r int = 1
	for y := n; y > k; y /= k {
		if y%k != 0 {
			fmt.Println(n, "is not a power of", k, ".")
			return
		}
		r++
	}
	fmt.Println(n, "=", k, "^", r)
}

// Test:
// Input: 64 4; 64 8; 9 2; 9 3
// Output: 64 = 4 ^ 3; 64 = 8 ^ 2; 9 is not a power of 2 .; 9 = 3 ^ 2
