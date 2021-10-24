package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is a perfect square.")
	var n float64
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	for i := 1; i < int(n); i++ {
		if int(n)/i == i {
			fmt.Printf("%v is a perfect square. %v^2 = %v", n, i, n)
			return
		}

	}
	fmt.Printf("%v is not a perfect square.", n)
}
