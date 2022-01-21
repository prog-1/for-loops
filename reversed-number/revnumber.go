package main

import "fmt"

func main() {
	fmt.Println("The program reverses the digits in the number.")
	fmt.Println("Enter the number:")
	var x uint
	var y uint
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	for rev := x; rev != 0; rev /= 10 {
		y = rev%10 + y*10
	}
	fmt.Println("reserved number", y)
}
