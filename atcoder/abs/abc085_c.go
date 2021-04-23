package main

import (
	"fmt"
)

func main() {
	var n, y int
	fmt.Scan(&n, &y)
	ans := []int{-1, -1, -1}
	for tenThousand := 0; tenThousand <= n; tenThousand++ {
		for fiveThousand := 0; fiveThousand <= n-tenThousand; fiveThousand++ {
			thousand := n - tenThousand - fiveThousand
			if tenThousand*10000+fiveThousand*5000+thousand*1000 == y {
				ans = []int{tenThousand, fiveThousand, thousand}
				break
			}
		}
	}
	fmt.Println(ans[0], ans[1], ans[2])
}
