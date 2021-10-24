package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is prime.")
	fmt.Println("Enter the number:")
	var a int
	fmt.Scanln(&a)
	var b int
	for b = 2; b < a; b++ {
		if a%b == 0 {
			fmt.Println("Is not a prime number", a)
			return
		}
	}
	fmt.Println("Is a prime number", a)
}