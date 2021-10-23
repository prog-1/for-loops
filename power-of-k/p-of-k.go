package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program determines if a number n is a power of k.")
	fmt.Println("Enter n and k:") // 64, 4; 64, 8; 9,2
	var n, k float64
	fmt.Scan(&n, &k)
	for i := 0.0; math.Pow(k, i) <= n; i++ {
		if math.Pow(k, i) == n {
			fmt.Println(n, "=", k, "^", i) // 64,4; 64,8
			return
		}
	}
	fmt.Println(n, "is not a power of", k) // 9, 2
}
