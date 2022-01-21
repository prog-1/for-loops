package main

import "fmt"

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	fmt.Println("Enter the number:")
	var i int
	fmt.Scanln(&i)
	var g, h int
	if i == 1 {
		fmt.Println("=2^0", i)
		return
	}
	if i <= 0 || i%2 != 0 {
		fmt.Println("Is not a power of 2", i)
		return
	}
	for r := i; r != 1; r /= 2 {
		g = r % 2
		h++
	}
	if g == 0 {
		fmt.Println(i, "=2^", h)
	} else {
		fmt.Println("Is not a power of 2", i)
	}
}
