package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	sr := []rune(s)
	var ans int
	for _, r := range sr {
		if r == '1' {
			ans += 1
		}
	}
	fmt.Println(ans)
}
