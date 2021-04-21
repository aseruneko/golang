package main

import (
	"fmt"
	"sort"
)

func main() {
	var l, n int
	fmt.Scan(&l, &n)
	var w []int
	for i := 0; i < n; i++ {
		w = append(w, 0)
		fmt.Scan(&w[i])
	}
	sort.Slice(w, func(i, j int) bool { return w[i] < w[j] })
	var num, sum int
	for i := 0; i < len(w); i++ {
		if sum+w[i] <= l {
			num += 1
			sum += w[i]
		}
	}
	fmt.Println(num)
}
