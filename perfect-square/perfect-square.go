package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is a perfect square.")
	var n, a int
	fmt.Println("Enter the number:")
	fmt.Scanln(&n)
	if n <= 0 {
		fmt.Print("error: ", n, " isn't a square")
		return
	}
	for a*a <= n {
		a++
		if a*a == n {
			fmt.Print(n, " is a perfect square.", n, "=", a, "^2")
			return
		}
	}
	fmt.Println(n, "is not a perfect square.")

}

// tested 16
// 16 is a perfect square.16=4^2
// tested 12
// 12 is not a perfect square.
// tested -4
// error: -4 isn't a square
