package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is a perfect square.")
	fmt.Println("Enter the number:")
	var n int
	i := 1
	a := 1
	fmt.Scan(&n)
	for i <= n {
		i++
		if i*i == n {
			a++
		}
	}
	if a == 2 {
		fmt.Println(n, "is a perfect square")
	} else {
		fmt.Println(n, "is not a perfect square.")
	}

}
