package main

import (
	"fmt"
)

func main() {
	dayOfMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var ans int
	for x := 1; x <= 12; x++ {
		for y := 1; y <= dayOfMonth[x-1]; y++ {
			sum := y/10 + y%10
			if x == sum {
				ans += 1
			}
		}
	}
	fmt.Println(ans)
}
