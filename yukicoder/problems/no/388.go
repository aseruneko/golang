package main

import (
	"fmt"
)

func main() {
	var s, f int
	fmt.Scan(&s, &f)
	var ans int = s/f + 1
	fmt.Println(ans)
}
