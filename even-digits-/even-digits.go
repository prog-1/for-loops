package main

import ( 
	"fmt"
)

func main() {
	fmt.Println("The program determines if all digits in the integer are even.")
	fmt.Println("Enter the number:")
	var a int
	fmt.Scanln(&a)
	var b int
	for evendigits := a; evendigits != 0; evendigits /= 10 {
		b = evendigits % 10
		if b%2 != 0 {
			fmt.Println("Not all digits in", n, "are even")
			return
		}
	}
	fmt.Println("All digits in", n, "are even")
}