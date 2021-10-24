package main

import ( 
	"fmt"
)

func main() {
	fmt.Println("The program determines if all digits in the integer are even.")
	fmt.Println("Enter the number:")
	var a int //-28; 379; 2806; 0; 33
	fmt.Scanln(&a)
	var b int
	for evendigits := a; evendigits != 0; evendigits /= 10 {
		b = evendigits % 10
		if b%2 != 0 {
			fmt.Println("Not all digits in", a, "are even") //379; 33
			return
		}
	}
	fmt.Println("All digits in", a, "are even") //-28; 2806; 0
}