package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, l int
	fmt.Scan(&x, &y, &l)
	var ans int
	if y > 0 {
		if y%l == 0 {
			ans += y / l
		} else {
			ans += y/l + 1
		}
	}
	if x != 0 {
		ans += 1
		if int(math.Abs(float64(x)))%l == 0 {
			ans += int(math.Abs(float64(x))) / l
		} else {
			ans += int(math.Abs(float64(x)))/l + 1
		}
		if y < 0 {
			ans += 1
		}
	} else {
		if y < 0 {
			ans += 2
		}
	}
	if y < 0 {
		if int(math.Abs(float64(y)))%l == 0 {
			ans += int(math.Abs(float64(y))) / l
		} else {
			ans += int(math.Abs(float64(y)))/l + 1
		}
	}
	fmt.Println(ans)
}
