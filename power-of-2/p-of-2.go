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

// Input: 1             Input: 2            Input: 0                       Input: 8
// Output: 1 = 2 ^ 0;   Output: 2 = 2 ^ 1   Output: 0 is not a power of 2  Output: 8 = 2 ^ 3
