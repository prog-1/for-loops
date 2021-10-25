package main

import "fmt"

func main() {
	fmt.Println("The program prints the first digit of the given integer.")
	fmt.Print("Enter the number: ")
	var x int
	fmt.Scan(&x)
	y := 0
	for ; x != 0; x /= 10 {
		y = y/10 + x%10
	}
	fmt.Println("The first digit is", y)
}
