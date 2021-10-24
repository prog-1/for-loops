package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is prime.")
	fmt.Println("Enter the number:")
	var n int
	a := 1
	i := 1
	fmt.Scan(&n)
	for i <= n {
		i++
		if n%i == 0 {
			a++
		}
	}
	if a == 2 {
		fmt.Println(n, "is a prime number")
	} else {
		fmt.Println(n, "is not a prime number")
	}

}
