package main

import (
	"fmt"
)

func main() {
	var k, n, f int
	fmt.Scan(&k, &n, &f)
	var a []int
	for i := 0; i < f; i++ {
		a = append(a, 0)
		fmt.Scan(&a[i])
	}
	var beans int = k * n
	for i := 0; i < f; i++ {
		beans -= a[i]
	}
	if beans < 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(beans)
	}
}
