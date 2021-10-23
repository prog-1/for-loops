package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is a perfect square.")
	var n, a int
	fmt.Println("Enter the number:")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Square does not exist")
		return
	}
	for x := n; a*a <= x; a++ {
		if a*a == x {
			fmt.Println(n, "is a perfect square. ", a, "^ 2 =", n)
			return
		}
	}
	fmt.Println(n, "is not a perfect square")
}
