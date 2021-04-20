package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var a []int
	for i := 0; i < n; i++ {
		a = append(a, 0)
		fmt.Scan(&a[i])
	}
	var sum int
	for _, num := range a {
		sum += num
	}
	fmt.Println(sum)
}
