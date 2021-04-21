package main

import (
	"fmt"
	"time"
)

func main() {
	var n, h, m, t int
	fmt.Scan(&n, &h, &m, &t)
	var now time.Time
	now = time.Date(2020, 1, 1, h, m, 0, 0, time.UTC)
	var dif int = (n - 1) * t
	dif = dif % 1440
	now = now.Add(time.Duration(dif) * time.Minute)
	var ansHour int = now.Hour()
	var ansMin int = now.Minute()
	fmt.Println(ansHour)
	fmt.Println(ansMin)
}
