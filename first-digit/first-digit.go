package main

import "fmt"

func main() {
	fmt.Println("The program prints the first digit of the given integer.")
	fmt.Print("Enter the number: ")
	var n int
	fmt.Scan(&n)
	y := 0
	for ; n != 0; n /= 10 {
		y = y/10 + n%10
	}
	fmt.Println("The first digit is", y)
} //input| 321 | 0 | 913 | 1 |
//output | 3   | 0 | 9   | 1 |
