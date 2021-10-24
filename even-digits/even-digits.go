package main

import "fmt"

func main() {
	fmt.Println("This program verifies whether a number is a palindrome.")
	fmt.Print("Enter a number: ")
	var x int
	fmt.Scan(&x)
	for tmp := x; tmp != 0; tmp /= 10 {
		if (tmp%10)%2 != 0 {
			fmt.Println("Not all digits in", x, "are even")
			return
		}
	}
	fmt.Println("All digits in", x, "are even")
}

// tested 42681
// Not all digits in 42681 are even
// tested 68426
// All digits in 68426 are even
// tested 1
// Not all digits in 1 are even
