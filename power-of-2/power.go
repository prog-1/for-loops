package main

import "fmt"

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	fmt.Print("Enter the number: ")
	var x int
	fmt.Scan(&x)
	if x%2 != 0 {
		fmt.Println(x, "is not a power of 2.")
		return
	} else if x == 0 {
		fmt.Println(x, "is not a power of 2.")
		return
	}
	var y int
	var tmp = x
	for ; tmp != 1; tmp /= 2 {
		if tmp%2 == 0 {
			y++
			continue
		} else {
			fmt.Println(x, "is not a power of 2.")
			break
		}
	}
	if tmp == 1 {
		fmt.Println(x, "=", "2 ^", y)
	}
}
