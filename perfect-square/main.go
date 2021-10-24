package main

import "fmt"

func main() {
	var num int
	fmt.Print("The program determines if an integer is a perfect square.\nEnter the number:")
	fmt.Scan(&num)
	a := 0
	for i := 1; i <= num; a = i * i {
		if a == num {
			fmt.Print(num, " is a perfect square.", i, "^2 = ", num)
			return
		}
		i++
	}
	fmt.Println(num, "is not a perfect square.")
}
