package main

import "fmt"

func main() {
	fmt.Println("The program reverses the digits in the number.")
	fmt.Println("Enter the number:")
	var n int
	fmt.Scanln(&n)
	var x int
	for ; n != 0; n /= 10 {
		x = x*10 + n%10
	}
	fmt.Println("The reversed number is", x)
}

// Test:
//
// Input: 0
// Output: The reversed number is 0
//
// Input: 23415
// Output: The reversed number is 51432
//
// Input: 102593165
// Output: The reversed number is 561395201
//
// Input: -4361
// Output: The reversed number is -1634
