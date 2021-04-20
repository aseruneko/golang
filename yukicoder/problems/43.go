package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var ans int
	switch b % a {
	case 0:
		ans = b / a
	default:
		ans = b/a + 1
	}
	fmt.Println(ans)
}
