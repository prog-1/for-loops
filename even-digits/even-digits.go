package main

import "fmt"

func main() {
	fmt.Println("The program determines if all digits in the integer are even.")
	fmt.Print("Enter the number: ")
	var n int
	fmt.Scan(&n)
	y := 0
	for tmp := n; tmp != 0; tmp /= 10 {
		y = y/10 + tmp%10
		if y%2 == 0 {
			continue
		} else {
			fmt.Println("Not all digits in", n, "are even.")
			return
		}
	}
	fmt.Println("All digits in", n, "are even.")
} //input|1              | 246788         | 3              |
//output |Not all digits | Not all digits | Not all digits |
