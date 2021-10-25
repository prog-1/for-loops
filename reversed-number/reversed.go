package main

import "fmt"

func main() {
	fmt.Println("The program reverses the digits in the number.")
	fmt.Print("Enter the number: ")
	var x int
	fmt.Scan(&x)
	y := 0
	for ; x != 0; x /= 10 {
		y = y*10 + x%10
	}
	fmt.Println("The reversed number is", y)
}
