package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program reverses the digits in the number.")
	fmt.Println("Enter the number:")
	var x uint
	fmt.Scan(&x)
	var y uint
	for z := x; z != 0; z /= 10 {
		y = y*10 + z%10
	}
	fmt.Println("The reversed number is", y)
}
