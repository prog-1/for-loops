package main

import "fmt"

func main() {
	fmt.Println("The program determines if an integer is prime.")
	fmt.Print("Enter the number: ")
	var n int
	fmt.Scan(&n)
	if n == 2 || n == 3 || n == 5 || n == 7 {
		fmt.Println(n, "is a prime number.")
	} else if n == 1 || n%2 == 0 || n%3 == 0 || n%5 == 0 || n%7 == 0 {
		fmt.Println(n, "is not a prime number.")
	} else {
		fmt.Println(n, "is a prime number.")
	}
} //input| 3        | 1      | 0      | 199      |
//output | is prime | is not | is not | is prime |
