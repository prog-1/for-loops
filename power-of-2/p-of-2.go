package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	fmt.Println("Enter the number:")
	var num float64 // 8; 16; 34; 64; 6
	fmt.Scan(&num)
	var i float64
	for i = 0; math.Pow(2, i) <= num; i++ {
		if math.Pow(2, i) == num {
			fmt.Println(num, "= 2 ^", i) // 8; 16; 64
			return
		}
	}
	fmt.Println(num, "is not a power of 2.") // 34; 6
}
