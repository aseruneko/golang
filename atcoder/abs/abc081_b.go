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
	var isDividable bool = true
	var ans int = -1
	for isDividable {
		ans += 1
		a, isDividable = div2(a)
	}
	fmt.Println(ans)
}

func div2(a []int) ([]int, bool) {
	var isDividable bool = true
	var output []int
	for i := range a {
		if a[i]%2 == 1 {
			isDividable = false
		}
		output = append(output, a[i]/2)
	}
	return output, isDividable
}
