package main

import "fmt"

func main() {
	fmt.Println("The program prints the first digit of the given integer.")
	fmt.Println("Enter the number:")
	var x int // 4457; -956; 234; 5
	fmt.Scan(&x)
	var y int
	for d := x; d != 0; d /= 10 {
		y = y/10 + d%10
	}
	fmt.Println("The first digit is", y) // 4; -9; 2; 5
}
