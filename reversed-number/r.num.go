package main

import "fmt"

func main() {
	fmt.Println("The program reverses the digits in the number.")
	fmt.Println("Enter the number:")
	var x int
	fmt.Scan(&x)
	var y int
	if x < 0 {
		fmt.Println("Number should be greater than 0 - palindrome can't be negative")
		return
	}
	for ; x != 0; x /= 10 {
		y = y*10 + x%10
	}
	fmt.Println("The reversed number is", y)
}

// Test:
// Input: 123, 1, 345, -5
// Output: The reversed number is 321, The reversed number is 1, The reversed number is 543, Number should be greater than 0 - palindrome can't be negative
