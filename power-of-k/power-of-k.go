package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a number n is a power of k.")
	var k, n int
	var i int = 0
	fmt.Print("Enter n and k: ")
	fmt.Scan(&n, &k)
	var d = n
	if n%k == 0 {
		for ; n != 1; n /= k {
			i++
		}
		fmt.Printf("%v = %v ^ %v", d, k, i)
		return
	}
	fmt.Printf("%v is not a power of %v.", n, k)
}
