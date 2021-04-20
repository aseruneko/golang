package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	var ans []rune
	fmt.Scan(&s)
	for _, r := range s {
		if unicode.IsUpper(r) {
			ans = append(ans, unicode.ToLower(r))
		} else {
			ans = append(ans, unicode.ToUpper(r))
		}
	}
	fmt.Println(string(ans))
}
