package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program reverses the digits in the number.")
	fmt.Println("Enter the number:")
	var a int
	fmt.Scan(&a)
	i := 0
	for a > 0 {
		i = (i * 10) + (a % 10)
		a = a / 10
	}
	fmt.Println("The reversed number is", i)
}
