package main

import "fmt"

func main() {
	fmt.Println("The program reverses the digits in the number.")
	fmt.Print("Enter the number: ")
	var n int
	var s = 0
	fmt.Scan(&n)
	for n > 0 {
		d := n % 10
		n = n / 10
		s = s*10 + d
	}
	fmt.Println("The reversed number is", s)
}
