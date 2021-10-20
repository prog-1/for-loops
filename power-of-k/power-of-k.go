package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a number n is a power of k.")
	var n, k int
	var i int = 0
	fmt.Print("Enter n and k: ")
	fmt.Scan(&n, &k)
	if n%k == 0 {
		for z := 1; z != n; z *= k {
			i++
			fmt.Println(i) // counter
		}
		fmt.Printf("%v = %v^%v", n, k, i)
	} else {
		fmt.Printf("%v is not a power of %v.", n, k)
	}
}
