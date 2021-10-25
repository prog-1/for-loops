package main

import "fmt"

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	fmt.Println("Enter the number:")
	var num int
	fmt.Scan(&num)
	var n int
	var res int
	for n = num; n > 0 && n%2 == 0; n /= 2 {
		res++
	}
	if n != 1 {
		fmt.Println(num, "is not a power of 2")
	} else {
		fmt.Println(num, "is a power of 2.", num, "= 2 ^", res)
	}
}

// Test (1)
// Input: 0
// Output: 0 is not a power of 2

// Test (2)
// Input: 1
// Output: 1 is a power of 2. 1 = 2 ^ 0

// Test (3)
// Input: 16
// Output: 16 is a power of 2. 16 = 2 ^ 4

// Test (4)
// Input: 362
// Output: 362 is not a power of 2

// Test (5)
// Input: -64
// Output: -64 is not a power of 2
