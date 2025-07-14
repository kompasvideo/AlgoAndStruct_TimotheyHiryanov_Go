package main

import "fmt"

func left_bound(a [10]int, n, x int) int {
	right := n
	left := -1
	for right-left > 1 {
		middle := (left + right) / 2
		if a[middle] < x {
			left = middle
		} else {
			right = middle
		}
	}
	return left
}
func find_binary(a [10]int, n, x int) int {
	left := left_bound(a, n, x)
	if left < n && a[left+1] == x {
		return left + 1
	}
	return -1
}
func main() {
	n := 10
	a := [10]int{1, 3, 3, 7, 7, 7, 7, 9, 10, 10}
	x := 9
	fmt.Println(a)
	fmt.Println("Left bound index ", left_bound(a, n, x))
	fmt.Println("Find index ", find_binary(a, n, x))
}
