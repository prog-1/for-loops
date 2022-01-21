package main

import "fmt"

func main() {
	fmt.Println("The program determines if all digits in the integer are even.")
	fmt.Println("Enter the number:")
	var x int
	fmt.Scanln(&x)
	var y int
	for edi := x; edi != 0; edi /= 10 {
		y = edi % 10
		if y%2 != 0 {
			fmt.Println("Digits not  in", x, "even")
			return
		}
	}
	fmt.Println("digits in", x, "even")
}
