package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a number n is a power of k.")
	fmt.Println("Enter n and k:")
	var n, k int64
	a := 0
	fmt.Scan(&n, &k)
	for n%k == 0 {
		n = n / k
		a++
	}
	if n == 1 {
		fmt.Println("your number =", k, "^", a)
	} else {
		fmt.Println("your number is not a power of ", k)
	}
}
