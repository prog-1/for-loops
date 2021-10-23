package main

import "fmt"

func main() {
	fmt.Println("The program determines if an integer is prime.")
	fmt.Println("Enter the number:")
	var num int //prime numbers are always possitive, because are >1
	fmt.Scan(&num)
	if num <= 1 {
		fmt.Println("Number must be greater than 1.")
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			fmt.Println(num, "is not a prime number")
			return
		}
	}
	fmt.Println(num, "is a prime number")
}
