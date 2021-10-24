package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a number is a power of k.")
	fmt.Println("Enter the number:")
	var n, k, a int
	fmt.Scanln(&n, &k)
	if n == 0 {
		fmt.Print("error: 0 isn't a power of any number")
	} else {
		var tmp = n
		if n >= k {
			for ; tmp != 1; tmp /= k {
				if tmp%k == 0 {
					a++
				} else {
					break
				}

			}
		}

		if tmp == 1 {
			fmt.Print(n, "=", k, "^", a)
		} else {
			fmt.Println(n, "is not a power of", k)
		}
	}

}

// tested 16 4
//16=4^2
//tested 0 5
//error: 0 isn't a power of any number
//tested 1
//1=2^0
//tested 5 3
//5 is not a power of 3
