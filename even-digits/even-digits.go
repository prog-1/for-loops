package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	var a, s float64
	if a == 0 {
		fmt.Print("error: 0 isn't a power of any number")
	} else {
		if a >= 2 {
			for fmt.Scan(&a); a > math.Pow(2, s); {
				s++
			}

		} else {
			for fmt.Scan(&a); a < math.Pow(2, s); {
				s--
			}
		}

		if math.Pow(2, s) == a {
			fmt.Print(a, "=2^", s)
		} else {
			fmt.Println(a, "is not a power of 2")
		}
	}
}

// tested 16
//16 is not a power of 2
//tested 9
//9 is not a power of 2
//tested 0.0625
//0.0625=2^-4
//tested 1
//1=2^0
//tested 0.06
//0.06 is not a power of 2
