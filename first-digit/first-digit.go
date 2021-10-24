package main

import "fmt"

func main() {
	fmt.Println("The program prints the first digit of the given integer.")
	fmt.Print("Enter the number: ")
	var n int
	var s = 0
	fmt.Scan(&n)
	for ; n != 0; n /= 10 {
		s = n % 10
	}
	fmt.Println("The first digit is", s)
	// fun fact if in number first is number, but then go letters program still works
	// but if be all letters result will be 0
}
