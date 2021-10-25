package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is prime.")
	var n, i int
	fmt.Println("enter a natural number:")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Error: prime numbers are the natural numbers")

	} else {

		if n == 1 {
			fmt.Print("1 is neither prime nor composite.")

		} else if n == 2 {
			fmt.Println("2 is a prime number.")

		} else {

			for i = 2; n-1 > i; i++ {
				if n%i == 0 {
					break

				}
			}
			if n%i == 0 {
				fmt.Println(n, "is not a prime number.")
			} else {
				fmt.Println(n, "is a prime number.")
			}

		}

	}

}

//tested 27
//27 is not a prime number.
//tested 2
//2 is a prime number.
//tested 0
//Error: prime numbers are the natural numbers
//tested 1
//1 is neither prime nor composite.
//tested 31
//31 is a prime number.
