package main

import "fmt"

func main() {
	fmt.Println("The program determines if a number n is a power of k.")
	fmt.Println("Enter n and k:")
	var n, k int
	fmt.Scanln(&n, &k)
	if n == 1 {
		fmt.Println(n, "=", k, "^ 0")
		return
	}
	if n <= 0 || n%k != 0 {
		fmt.Println(n, "is not a power of", k)
		return
	}
	var x, y int
	for n2 := n; n2 != 1; n2 /= k {
		x = n2 % k
		y++
	}
	if x == 0 {
		fmt.Println(n, "=", k, "^", y)
	} else {
		fmt.Println(n, "is not a power of", k)
	}
}

// Test:
//
// Input: 0 2
// Output: 0 is not a power of 2
//
// Input: 1 4
// Output: 1 = 4 ^ 0
//
// Input: 16 3
// Output: 16 is not a power of 3
//
// Input: 16 4
// Output: 16 = 4 ^ 2
//
// Input: -8 2
// Output: -8 is not a power of 2
