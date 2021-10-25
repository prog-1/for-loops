package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	if n%2 == 0 && n != 2 || n%3 == 0 && n != 3 || n%5 == 0 && n != 5 || n%7 == 0 && n != 7 {
		fmt.Println(n, "is not a prime number.")
	} else {
		fmt.Println(n, "is a prime number.")
	}
}
