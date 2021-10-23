package main

import "fmt"

func main() {
	fmt.Println("The program determines if a number n is a power of k.")
	fmt.Print("Enter n and k: ")
	var n, k uint
	fmt.Scan(&n, &k)
	if k == 0 {
		fmt.Println("Any number raised to the power 0 will be equal to 1")
		return
	} else if n%k != 0 {
		fmt.Println(n, "is not a power of", k)
		return
	} else if n == 0 {
		fmt.Println(n, "=", n, "^", k)
		return
	}
	var a int
	var tmp = n
	for ; tmp != 1; tmp /= k {
		if tmp%k == 0 {
			a++
			continue
		} else {
			fmt.Println(n, "is not a power of", k)
			break
		}
	}
	if tmp == 1 {
		fmt.Println(n, "=", k, "^", a)
	}
} //input| 3 6    | 6 2    | 3 0     | 0 10   | 128 2 |
//output | is not | is not | line 11 | 0 ^ 10 | 2 ^ 7 |
