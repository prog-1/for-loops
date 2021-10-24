package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program prints the first digit of the given integer.")
	fmt.Print("Enter the number:")
	var a int
	fmt.Scan(&a)
	i := 10
	f := 0
	for ; a != 0; a /= 10 {
		f = a % 10

	}
	i += 10 * 10
	fmt.Println("The first digit is:", f)

}
