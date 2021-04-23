package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var sum int = a * b
	if sum%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
