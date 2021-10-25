package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	a := 0
	for i := 0; n > 0; n /= 10 {
		i = n % 10
		a = a*10 + i
	}
	fmt.Println("The reversed number is", a)
}
