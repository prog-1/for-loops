package main

import "fmt"

func main() {
	fmt.Println("The program determines if all digits in the integer are even.")
	fmt.Print("Enter the number: ")
	var x int
	fmt.Scan(&x)
	y := 0
	for tmp := x; tmp != 0; tmp /= 10 {
		y = y/10 + tmp%10
		if y%2 == 0 {
			continue
		} else {
			fmt.Println("Not all digits in", x, "are even.")
			return
		}
	}
	fmt.Println("All digits in", x, "are even.")
}
