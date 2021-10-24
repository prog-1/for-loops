package main

import "fmt"

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	fmt.Println("Enter the number:")
	var x int
	fmt.Scan(&x)
	if x == 1 {
		fmt.Println(x, "=2^0")
		return
	}
	if x%2 != 0 {
		fmt.Println(x, "is not a power of 2.")
		return
	}
	var r int = 1
	for y := x; y > 2; y /= 2 {
		if y%2 != 0 {
			fmt.Println(x, "is not a power of 2")
			return
		}
		r++
	}
	fmt.Println(x, "=2^", r)
}

// Test:
// Input: 8, 6, 67, 1
// Output: 2 =2^ 3; 6 is not a power of 2; 67 is not a power of 2.; 1 =2^0
