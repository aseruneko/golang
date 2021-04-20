package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		if strings.Contains(strconv.Itoa(i), "3") || i%3 == 0 {
			fmt.Println(i)
		}
	}
}
