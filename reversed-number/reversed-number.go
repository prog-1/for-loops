package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program reverses the digits in the number.")
	fmt.Println("Enter the number:")
	var d uint
	fmt.Scanln(&d)
	var e uint
	for rnumber := d; rnumber != 0; rnumber /= 10 {
		e = e*10 + rnumber%10
	}
	fmt.Println("The reversed number is", e)
}