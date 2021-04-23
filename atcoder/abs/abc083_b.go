package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	var ans int
	for i := 1; i <= n; i++ {
		sum := i/10000 + (i/1000)%10 + (i/100)%10 + (i/10)%10 + i%10
		if sum >= a && sum <= b {
			ans += i
		}
	}
	fmt.Println(ans)
}
