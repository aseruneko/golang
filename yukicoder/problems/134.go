package main

import (
	"fmt"
)

func main() {
	var l, k int
	fmt.Scan(&l, &k)
	var ans int = (l - 1) / (k * 2) * k
	fmt.Println(ans)
}
