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
		if k == 0 {
			fmt.Println("power is any number except 0")
			return
		} else {
			fmt.Print(n, "is not a power of", k)
			return
		}

	}
	var tmp = n
	if n > 0 {

		for ; tmp != 1; tmp /= k {
			if tmp%k == 0 {
				a++
			} else {
				break
			}

		}

	} else {
		for ; tmp != -1; tmp /= k {
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

// tested 16 4
//16=4^2
//tested 0 0
//power is any number except 0
//tested 1
//1=2^0
//tested 5 3
//5 is not a power of 3
//tested -64 4
//-64 is not a power of 4
// tested -8 -2
// -8=-2^3
