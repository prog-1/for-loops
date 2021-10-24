package main

import "fmt"

func main() {
	fmt.Println("This program verifies whether a number is a palindrome.")
	fmt.Print("Enter a number: ")
	var x int
	fmt.Scan(&x)
	var y int
	for tmp := x; tmp != 0; tmp /= 10 {
		y = y*10 + tmp%10
	}
	fmt.Println("The reversed number is", y)
}

//tested 312
//The reversed number is 213
//tested -5215
//The reversed number is -5125
//tested 2348126471246921
//The reversed number is 1296421746218432
