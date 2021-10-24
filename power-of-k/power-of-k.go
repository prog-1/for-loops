package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a number n is a power of k.")
	fmt.Println("Enter n and k:")
	var n, k int //64, 4; 8, 3; 81, 9; 3, 2
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
	for r := n; r != 1; r /= k {
		a = r % k
		b++
	}
	if a == 0 {
		fmt.Println(n, "=", k, "^", b) //64, 4; 81, 9
	} else {
		fmt.Println(n, "Is not a power of", k) //8, 3; 3, 2
	}
}