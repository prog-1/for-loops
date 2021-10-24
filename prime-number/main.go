package main

import "fmt"

func main() {
	var num int
	fmt.Print("The program determines if an integer is prime.\nEnter the number: ")
	fmt.Scan(&num)
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			fmt.Println(num, "is not a prime number.")
			return
		}
	}
	if num == 1 {
		fmt.Println(num, "is not a prime number.")
	}
	fmt.Println(num, "is a prime number.")
}
