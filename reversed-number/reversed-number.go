package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program reverses the digits in the number.")
	fmt.Println("Enter the number:")
	var d uint //0; 2; 346; 4587
	fmt.Scanln(&d)
	var e uint
	for rnumber := d; rnumber != 0; rnumber /= 10 {
		e = e*10 + rnumber%10
	}
	fmt.Println("The reversed number is", e) //0; 2; 643; 7854
}