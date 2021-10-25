package main

import "fmt"

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	fmt.Println("Enter the number:")
	var num int
	fmt.Scan(&num)
	var n int
	var res int
	for n = num; n > 0 && n%2 == 0; n /= 2 {
		res++
	}
	if n != 1 {
		fmt.Println(num, "is not a power of 2")
	} else {
		fmt.Println(num, "is a power of 2.", num, "= 2 ^", res)
	}
}
