package main

import "fmt"

func find(a [7]int, n, x int) int {
	for i := 0; i < n; i++ {
		if a[i] == x {
			return i
		}
	}
	return -1
}
func main() {
	n := 7
	a := [7]int{5, 2, 3, 4, 6, 2, 3}
	fmt.Println(a)
	fmt.Println(find(a, n, 6))
}
