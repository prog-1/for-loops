package main

import "fmt"

func main() {
	fmt.Println("The program determines if all digits in the integer are even.")
	fmt.Println("Enter the number:")
	var x int // 456; 21; -5321; 2286; 24
	fmt.Scan(&x)
	for ed := x; ed != 0; ed /= 10 {
		if (ed%10)%2 != 0 {
			fmt.Println("Not digits in", x, "are even.") // 456; 21; -5321
			return
		}
	}
	fmt.Println("All digits in", x, "are even.") // 2286; 24
}
