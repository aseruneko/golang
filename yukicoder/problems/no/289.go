package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	var sum int
	for _, num := range s {
		if unicode.IsDigit(num) {
			str, _ := strconv.Atoi(string(num))
			sum += str
		}
	}
	fmt.Println(sum)
}
