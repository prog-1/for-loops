package main

import "fmt"

func main() {
	fmt.Println("The program determines if a number n is a power of k.")
	fmt.Print("Enter n and k: ")
	var num, k int
	fmt.Scan(&num)
	fmt.Scan(&k)
	i := 1
	for result := k; result <= num; result *= k {
		if result == num {
			fmt.Print(num, " = ", k, "^", i, "\n")
			return
		}
		i += 1
	}
	fmt.Println(num, "is not a power of", k)
}
