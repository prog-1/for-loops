package main

import "fmt"

func main() {
	var n float64
	var k float64
	fmt.Print("Enter n and k: ")
	fmt.Scan(&n, &k)
	a := n
	power := 1
	if n == 1 && k == 1 {
		fmt.Println(n, "=", k, "^ 0")
	} else {
		for ; a != 1; power++ {
			a = a / k
			if a < k {
				break
			}
		}
		if a == 1 {
			fmt.Println(n, "=", k, "^", power)
		} else {
			fmt.Println(n, "is not a power of", k)
		}
	}
}
