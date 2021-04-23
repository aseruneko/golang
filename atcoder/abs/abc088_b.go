package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	var a []int
	for i := 0; i < n; i++ {
		a = append(a, 0)
		fmt.Scan(&a[i])
	}
	var alice, bob int
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	for i := range a {
		if i%2 == 0 {
			alice += a[i]
		} else {
			bob += a[i]
		}
	}
	ans := alice - bob
	fmt.Println(ans)
}
