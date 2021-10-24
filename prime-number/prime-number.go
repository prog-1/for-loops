package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is prime.")
	fmt.Print("Enter the number:")
	var a int
	fmt.Scan(&a)
	for i := 2; i < a; i++ {
		if a%i == 0 {
			fmt.Print(a, " is not a prime number.")
			return

		}

	}
	fmt.Print(a, " is a prime number.")

}
