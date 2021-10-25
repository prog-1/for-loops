package main

import "fmt"

func main() {
	fmt.Println("The program determines if an integer is prime.")
	fmt.Println("Enter the number:")
	var num int
	fmt.Scan(&num)
	if num <= 1 {
		fmt.Println("Number must be greater than 1. It is not a prime number") //prime numbers are always possitive, because are >1
		return
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			fmt.Println(num, "is not a prime number")
			return
		}
	}
	fmt.Println(num, "is a prime number")
}

// Test(1)
// Input: 0
// Output: Number must be greater than 1. It is not a prime number

// Test(2)
// Input: 1
// Output: Number must be greater than 1. It is not a prime number

// Test(3)
// Input: 2
// Output: 2 is a prime number

// Test(4)
// Input: 7
// Output: 7 is a prime number

// Test(5)
// Input: 34
// Output: 34 is not a prime number

// Test(6)
// Input: 482
// Output: 482 is not a prime number

// Test(7)
// Input: -3
// Output: Number must be greater than 1. It is not a prime number
