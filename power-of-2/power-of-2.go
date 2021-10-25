package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	fmt.Println("Enter the number:")
	var n, k int
	fmt.Scanln(&n)
	if n == 0 {
		fmt.Print(n, "isn't a power of 2")
	} else {
		var tmp = n
		if n >= 2 {
			for ; tmp != 1; tmp /= 2 {
				if tmp%2 == 0 {
					k++
				} else {
					break
				}

			}
		}

		if tmp == 1 {
			fmt.Print(n, "=2^", k)
		} else {
			fmt.Println(n, "is not a power of 2")
		}
	}

}

// tested 16
//16=2^4
//tested 9
//9 is not a power of 2
//tested 1
//1=2^0
//tested 0.06
//0.06 is not a power of 2
//tested -8
//-8 is not a power of 2
//tested 0
//error: 0 isn't a power of any number
