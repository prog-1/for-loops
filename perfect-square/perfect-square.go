package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program determines if an integer is a perfect square.")
	fmt.Print("Enter the number: ")
	var n int
	fmt.Scan(&n)
	var tmp float64 = float64(n)
	tmp = math.Sqrt(tmp)
	var a int = int(tmp)
	if a*a == n {
		fmt.Println(n, "is a perfect square.", a, "^ 2 =", n)
	} else {
		fmt.Println(n, "is not a perfect square.")
	}
} //input| 9     | 1     | 0     | 56     | 4096   |
//output | 3 ^ 2 | 1 ^ 2 | 0 ^ 2 | is not | 64 ^ 2 |
