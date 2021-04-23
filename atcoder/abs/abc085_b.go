package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var d []int
	for i := 0; i < n; i++ {
		d = append(d, 0)
		fmt.Scan(&d[i])
	}
	uniqued := Of(d).ToSlice()
	ans := len(uniqued)
	fmt.Println(ans)
}

type IntSet struct {
	value []int
}

func Of(slice []int) *IntSet {
	mapSlice := make(map[int]bool)
	var uniqueSlice []int
	for _, element := range slice {
		if !mapSlice[element] {
			mapSlice[element] = true
			uniqueSlice = append(uniqueSlice, element)
		}
	}
	return &IntSet{value: uniqueSlice}
}

func (t IntSet) ToSlice() []int {
	return t.value
}
