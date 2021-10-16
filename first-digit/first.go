package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the sum of digits of N.")
	var a int
	fmt.Scanln(&a)
	if a >= 0 {
		for ; a >= 10; a /= 10 {
		}
		s := a % 10
		fmt.Println("the first digit is", s)
	} else {
		for ; a <= -10; a /= 10 {
		}
		s := a % 10
		fmt.Println("the first digit is", math.Abs(float64(s)))
	}
}

// tested -87321
// the first digit is 8
// tested54372
// the first digit is 5
// tested 0
// the first digit is 0
