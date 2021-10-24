package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a number n is a power of k.")
	fmt.Println("Enter n and k:")
	var n, k int
	fmt.Scanln(&n, &k)
	var a, b int
	if n == 1 {
		fmt.Println(n, "=", b, "^0")
		return
	}
	if n < 0 || n%k != 0 {
		fmt.Println(n, "Is not a power of", k)
		return
	}
	for c := n; x != 1; x /= k {
		a = c % k
		b++
	}
	if a == 0 {
		fmt.Println(n, "=", k, "^", b)
	} else {
		fmt.Println(n, "Is not a power of", k)
	}
}