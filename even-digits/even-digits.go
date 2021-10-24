package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if all digits in the integer are even.")
	fmt.Print("Enter the number:")
	var a int
	fmt.Scan(&a)
	r := a
	for ; a != 0; a /= 10 {
		if (a%10)%2 != 0 {
			fmt.Print("Not all digits in ", r, " are even.")
			return

		}

	}
	fmt.Print("All digits in ", r, " are even.")

}
