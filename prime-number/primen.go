package main

import "fmt"

func main() {
	fmt.Println("The program determines if an integer is prime.")
	fmt.Println("Enter the number:")
	var number int
	fmt.Scanln(&number)
	var x int
	for x = 2; x < number; x++ {
		if number%x == 0 {
			fmt.Println("not a prime number")
			return
		}
	}
	fmt.Println("is a prime number", number)
}
