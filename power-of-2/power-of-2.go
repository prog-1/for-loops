package main

import "fmt"

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	var n, i int
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	if n > 0 && (n&(n-1)) == 0 {
		for z := 1; z != n; z *= 2 {
			i++
		}
		fmt.Printf("%v = 2^%v", n, i)
	} else {
		fmt.Println(n, "is not a power of 2.")
	}

}
