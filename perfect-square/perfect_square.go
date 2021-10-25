package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	var k int
	for ; k*k < n; k++ {
	}
	if k*k > n {
		fmt.Println(n, "is not a perfect square.")
	} else {
		fmt.Println(n, "is a perfect square.", k, "^ 2 =", n)
	}
}
