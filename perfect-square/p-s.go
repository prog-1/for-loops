package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is a perfect square.")
	fmt.Println("Enter the number:")
	var x int // 9; 16; 8; 23; 45
	fmt.Scan(&x)
	for i := 0; i*i <= x; i++ {
		if i*i == x {
			fmt.Println(x, "is a perfect square.", i, "^2=", x) // 9; 16
			return
		}
	}
	fmt.Println(x, "is not a perfect square.") // 8; 23; 45
}
