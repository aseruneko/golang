package main

import (
	"fmt"
)

func main() {
	var l, m, n int
	fmt.Scan(&l, &m, &n)
	m += n / 25
	n = n % 25
	l += m / 4
	m = m % 4
	l = l % 10
	var ans int = l + m + n
	fmt.Println(ans)
}
