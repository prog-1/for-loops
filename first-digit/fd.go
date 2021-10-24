package main

import "fmt"

func main() {
	fmt.Println("The program prints the first digit of the given integer.")
	fmt.Println("Enter the number:")
	var x int
	fmt.Scan(&x)
	if x > 10 {
		for x >= 10 {
			x /= 10
		}
		fmt.Println("The first digit is", x)
	} else {
		for x <= -10 {
			x /= 10
		}
		fmt.Println("The first digit is", -x)
	}
}

// Test:
// Input: 1532, -34, -245, 0, 2147
// Output: The first digit is 1, The first digit is 3, The first digit is 2, The first digit is 3, The first digit is 0, The first digit is 2
