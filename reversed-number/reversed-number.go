package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program reverses the digits in the number.")
	fmt.Print("Enter the number:")
	var a int
	fmt.Scan(&a)
	r := 0
	for ; a != 0; a /= 10 {
		r = a%10 + (r * 10)
	}
	fmt.Print("The reversed number is: ", r)

}
