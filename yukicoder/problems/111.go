package main

import (
	"fmt"
	"math"
)

func main() {
	var d, p int
	fmt.Scan(&d, &p)
	var ans int = int(math.Floor(float64(d*100+d*p) / 100))
	fmt.Println(ans)
}
