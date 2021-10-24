package main

import "fmt"

func main() {
	fmt.Print("The program prints the first digit of the given integer.\nEnter the number:")
	var num int
	fmt.Scan(&num)
	var res int
	for ; num != 0; num /= 10 {
		res = num % 10
	}
	fmt.Print("The first digit is ", res)
}
