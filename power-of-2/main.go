package main

import "fmt"

func main() {
	fmt.Println("The program determines if a number is a power of 2.")
	fmt.Print("Enter the number: ")
	var num int
	fmt.Scan(&num)
	i := 1
	for result := 2; result <= num; result *= 2 {
		if result == num {
			fmt.Print(num, " = 2^", i, "\n")
			return
		}
		i += 1
	}
	fmt.Println(num, "is not a power of 2")
}
