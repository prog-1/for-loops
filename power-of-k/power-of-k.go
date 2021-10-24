package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a number n is a power of k.")
	fmt.Print("Enter n and k:")
	var n, k int
	fmt.Scan(&n, &k)
	var i int
	result := k
	i = 1
	for {
		if n == result {
			fmt.Print(k, "^", i, " == ", n)
			return

		}
		if n < result {
			fmt.Print(n, " is not a power of ", k)
			return

		}
		i += 1
		result *= k
	}

}
