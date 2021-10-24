package main

import "fmt"

func main() {
	var num int
	fmt.Print("The program determines if all digits in the integer are even.\nEnter the number: ")
	fmt.Scan(&num)
	tmp := num
	for ; num != 0; num /= 10 {
		if (num%10)%2 != 0 {
			fmt.Print("Not all digits in ", tmp, " are even.")
			return
		}
	}
	fmt.Print("All digits in ", tmp, " are even.")
}
