package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var ans string
	switch (n + 3 - k) % 3 {
	case 0:
		ans = "Drew"
	case 2:
		ans = "Won"
	case 1:
		ans = "Lost"
	}
	fmt.Println(ans)
}
