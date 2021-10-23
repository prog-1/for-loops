package main

import "fmt"

func main() {
	fmt.Println("The program reverses the digits in the number.")
	fmt.Println("Enter the number:")
	var x uint // 765; 389; -87; 95849
	fmt.Scan(&x)
	var y uint
	for rn := x; rn != 0; rn /= 10 {
		y = y*10 + rn%10
	}
	fmt.Println("The reversed number is", y) // 567; 983; 0 (because "uint" only for possitive numbers); 94859
}
