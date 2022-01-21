package main

import "fmt"

func main() {
	fmt.Println("The program prints the first digit of the given integer.")
	fmt.Println("Enter the number:")
	var x int
	fmt.Scanln(&x)
	var y int
	for y = x; y/10 != 0; {
		y /= 10
	}
	fmt.Println("First digit", "y")

}
