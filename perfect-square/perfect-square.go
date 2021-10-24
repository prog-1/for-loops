package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if an integer is a perfect square.")
	fmt.Println("Enter the number:") 
	var n int //4; 11; 23; 25
	fmt.Scanln(&n)
	for i := 0; i*i <= n; i++ {
		if i*i == n {
			fmt.Println(n, "Is a perfect square", i, "^2=", n) //4; 25
			return
		}
	}
	fmt.Println("Is not a perfect square", n) //11; 23
}