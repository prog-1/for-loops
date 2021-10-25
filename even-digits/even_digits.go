package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	i := 0
	a := n
	for ; a > 0; a /= 10 {
		i = a % 10
		if i%2 != 0 {
			fmt.Println("Not all digits in", n, "are even.")
			break
		} else if i%2 == 0 && a < 10 {
			fmt.Println("All digits in", n, "are even.")
		}
	}
}
