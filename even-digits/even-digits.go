package main

import "fmt"

func main() {
	fmt.Println("The program determines if all digits in the integer are even.")
	fmt.Print("Enter the number: ")
	var n, s int
	fmt.Scan(&n)
	for ; n != 0; n /= 10 {
		d := n
		s = d % 10
		if s%2 != 0 {
			fmt.Printf("Not all digits in %v are even.", n)
			break
		} else {
			fmt.Printf("All digits in %v are even.", n)
			break
		}

	}
}
