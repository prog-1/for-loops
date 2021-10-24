package main

import "fmt"

func main() {
	var n int
	fmt.Println("The program determines if an integer is prime.")
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	for i := 1; i != n+1; i++ {
		if n%2 == 0 && n != 2 ||
			n%3 == 0 && n != 3 ||
			n%5 == 0 && n != 5 ||
			n%7 == 0 && n != 7 {
			fmt.Printf("%v is not a prime number.", n)
			break
		} else {
			fmt.Printf("%v is a prime number.", n)
			break
		}
	}
}
