package main

import "fmt"

func main() {
	var n float64
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	a := n
	i := 1
	if a == 1 {
		fmt.Println(n, "= 2 ^ 0")
	} else {
		for ; a != 1; i++ {
			a = a / 2
			if a < 2 {
				break
			}
		}
		if a == 1 {
			fmt.Println(n, "= 2 ^", i)
		} else {
			fmt.Println(n, "is not a power of 2.")
		}
	}
}
