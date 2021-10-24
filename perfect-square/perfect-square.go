package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is a perfect square.")
	fmt.Println("Enter the number:")
	var n int
	fmt.Scanln(&n)
	for i := 0; i*i <= n; i++ {
		if i*i == n {
			fmt.Println(n, "Is a perfect square", i, "^2=", n)
			return
		}
	}
	fmt.Println("Is not a perfect quare", x)
}