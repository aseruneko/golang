package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var ans int
	for i := 0; int(math.Pow(float64(2), float64(i))) <= n; i++ {
		ans = i
	}
	if n-int(math.Pow(float64(2), float64(ans))) != 0 {
		ans += 1
	}
	fmt.Println(ans)
}
