package main

import "fmt"

func main() {
	var n int
	fmt.Println("The program determines if an integer is prime.")
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	d := n
	for s := 2; s != d; s++ {
		if n%s == 0 {
			fmt.Printf("%v is not a prime number.", n)
			return
		}
	}
	fmt.Printf("%v is a prime number.", n)
}
