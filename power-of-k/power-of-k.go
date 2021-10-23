package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program determines if a number n is a power of k.")
	var k int
	var i int = 0
	var n float64
	fmt.Print("Enter n and k: ")
	fmt.Scan(&n, &k)
	if n/(math.Sqrt(n)) == (math.Sqrt(n)) {
		for z := 1; z != int(n); z *= k {
			i++ // counter

		}
		fmt.Printf("%v = %v^%v", n, k, i)
	} else {
		fmt.Printf("%v is not a power of %v.", n, k) // why this line does not work???
	}
}
