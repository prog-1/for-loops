package main

import "fmt"

func main() {
	fmt.Println("The program determines if a number n is a power of k.")
	fmt.Print("Enter n and k: ")
	var x, y uint
	fmt.Scan(&x, &y)
	if y == 0 {
		fmt.Println("Any number raised to the power 0 will be equal to 1")
		return
	} else if x%y != 0 {
		fmt.Println(x, "is not a power of", y)
		return
	} else if x == 0 {
		fmt.Println(x, "=", x, "^", y)
		return
	}
	var z int
	var tmp = x
	for ; tmp != 1; tmp /= y {
		if tmp%y == 0 {
			z++
			continue
		} else {
			fmt.Println(x, "is not a power of", y)
			break
		}
	}
	if tmp == 1 {
		fmt.Println(x, "=", y, "^", z)
	}
}
