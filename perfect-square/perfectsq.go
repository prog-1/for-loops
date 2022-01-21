package main

import "fmt"

func main() {
	fmt.Println("The program determines if an integer is prime.")
	fmt.Println("Enter the number:")
	var x int
	fmt.Scanln(&x)
	for y := 0; y*y <= x; y++ {
		if x == y*y {
			fmt.Println("perfet square", x, y, "^2=")
			return
		}
	}
	fmt.Println("Not perfect square", "x")
}
