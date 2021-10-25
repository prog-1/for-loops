package main

import "fmt"

func main() {
	fmt.Println("The program determines if an integer is prime.")
	fmt.Print("Enter the number: ")
	var x int
	fmt.Scan(&x)
	if x == 2 || x == 3 || x == 5 || x == 7 {
		fmt.Println(x, "is a prime number.")
	} else if x == 1 || x%2 == 0 || x%3 == 0 || x%5 == 0 || x%7 == 0 {
		fmt.Println(x, "is not a prime number.")
	} else {
		fmt.Println(x, "is a prime number.")
	}
}
