package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	var s string
	fmt.Scan(&a, &b, &s)
	var ans string = strconv.Itoa(a+b) + " " + s
	fmt.Println(ans)
}
