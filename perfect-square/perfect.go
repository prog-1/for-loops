package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program determines if an integer is a perfect square.")
	fmt.Print("Enter the number:")
	var x int
	fmt.Scan(&x)
	var tmp float64 = float64(x)
	tmp = math.Sqrt(tmp)
	var y int = int(tmp)
	if y*y == x {
		fmt.Println(x, "is a perfect square.", y, "^ 2 =", x)
	} else {
		fmt.Println(x, "is not a perfect square.")
	}
}
