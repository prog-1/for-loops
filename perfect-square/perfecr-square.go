package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is a perfect square.")
	fmt.Print("Enter the number:")
	var a int
	fmt.Scan(&a)
	i := 1
	for {
		if i*i == a {
			fmt.Print(a, " is a perfect square.", i, "^2 = ", a)
			return
		}
		if i*i > a {
			fmt.Print(a, " is not a perfect square.")
			return
		}
		i += 1

	}
}
