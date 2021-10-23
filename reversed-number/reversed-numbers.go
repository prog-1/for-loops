package main

import "fmt"

func main() {
	fmt.Println("The program reverses the digits in the number.")
	fmt.Print("Enter the number: ")
	var n int
	fmt.Scan(&n)
	y := 0
	for ; n != 0; n /= 10 {
		y = y*10 + n%10
	}
	fmt.Println("The reversed number is", y)
} //input| 431 | 192743 | 019                     | 19 |
//output | 134 | 347291 | 1 didn't understand why | 91 |
