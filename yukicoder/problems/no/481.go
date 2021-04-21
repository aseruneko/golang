package main

import "fmt"

func main() {
	correct := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 9; i++ {
		fmt.Scan(&b[i])
	}
	for i := 0; i < 9; i++ {
		if correct[i] != b[i] {
			fmt.Println(correct[i])
			break
		}
		if i == 8 {
			fmt.Println(10)
		}
	}
}
