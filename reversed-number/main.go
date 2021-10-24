package main

import "fmt"

func main() {
	fmt.Print("The program reverses the digits in the number.\nEnter the number:")
	var num int
	fmt.Scan(&num)
	var res int
	for ; num != 0; num /= 10 {
		res = res*10 + num%10
	}
	fmt.Print("The reversed number is ", res)
}
