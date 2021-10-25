package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program prints the first digit of the given integer.")
	fmt.Println("Enter number:")
	var a int
	fmt.Scan(&a)
	for a >= 10 {
		a = a / 10
	}
	fmt.Println(a)

}
