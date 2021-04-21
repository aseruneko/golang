package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	ar := []rune(a)
	br := []rune(b)
	sort.Slice(ar, func(i, j int) bool { return ar[i] > ar[j] })
	sort.Slice(br, func(i, j int) bool { return br[i] > br[j] })
	ars := string(ar)
	brs := string(br)
	if ars == brs {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
