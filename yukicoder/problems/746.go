package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var a []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		a = append(a, 0)
		fmt.Scan(&a[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	var ans float64
	switch len(a) % 2 {
	case 1:
		ans = float64(a[len(a)/2])
	case 0:
		ans = (float64(a[len(a)/2-1]) + float64(a[len(a)/2])) / 2
	}
	fmt.Println(ans)
}
