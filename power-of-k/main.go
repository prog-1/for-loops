package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a number n is a power of k.")
	fmt.Println("Enter n and k:")
	var n, k int
	fmt.Scan(&n, &k)
	var y, z int
	if n == 1 {
		fmt.Println(n, "=", z, "^ 0")
		return
	}
	if n < 0 || n%k != 0 {
		fmt.Println(n, "is not a power of", k)
		return
	}
	for x := n; x != 1; x /= k {
		y = x % k
		z++
	}
	if y == 0 {
		fmt.Println(n, "=", k, "^", z)
	} else {
		fmt.Println(n, "is not a power of", k)
	}
}
