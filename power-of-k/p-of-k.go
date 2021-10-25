package main

import "fmt"

func main() {
	fmt.Println("The program determines if a number n is a power of k.")
	fmt.Println("Enter n and k:")
	var n, k int
	fmt.Scan(&n, &k)
	var num int
	var res int
	for num = n; num > 0 && num%k == 0; num /= k {
		res++
	}
	if num != 1 {
		fmt.Println(n, "is not a power of", k)
	} else {
		fmt.Println(n, "=", k, "^", res)
	}
}

// Test (1)
// Input: 1 3
// Output: 1 = 3 ^ 0

// Test (2)
// Input: 9 2
// Output: 9 is not a power of 2

// Test (3)
// Input: 64 8
// Output: 64 = 8 ^ 2

// Test (4)
// Input: -3 2
// Output: -3 is not a power of 2

// Test (5)
// Input: 8 -2
// Output: 8 is not a power of -2

// Test (6)
// Input: 8 2
// Output: 8 = 2 ^ 3
