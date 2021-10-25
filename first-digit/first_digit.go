package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	for n > 10 {
		n = n / 10
	}
	fmt.Println("The first digit is", n)
}
