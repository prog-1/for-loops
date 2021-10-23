package main

import "fmt"

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	fmt.Print("Enter the number: ")
	var n int
	fmt.Scan(&n)
	if n%2 != 0 {
		fmt.Println(n, "is not a power of 2.")
		return
	} else if n == 0 {
		fmt.Println(n, "is not a power of 2.")
		return
	}
	var a int
	var tmp = n
	for ; tmp != 1; tmp /= 2 {
		if tmp%2 == 0 {
			a++
			continue
		} else {
			fmt.Println(n, "is not a power of 2.")
			break
		}
	}
	if tmp == 1 {
		fmt.Println(n, "=", "2 ^", a)
	}
} //input| 8     | 16    | 0      | 1      | 2048   |
//output | 2 ^ 3 | 2 ^ 4 | is not | is not | 2 ^ 11 |
