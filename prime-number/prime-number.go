package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is prime.")
	fmt.Println("Enter the number:")
	var a int
	fmt.Scanln(&a) //3; 13; 24; 68
	var b int
	for b = 2; b < a; b++ {
		if a%b == 0 {
			fmt.Println("Is not a prime number", a) //24; 68
			return
		}
	}
	fmt.Println("Is a prime number", a) //3; 13
}