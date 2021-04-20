package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var ans int = (1 + n) * n / 2
	fmt.Println(ans)
}
