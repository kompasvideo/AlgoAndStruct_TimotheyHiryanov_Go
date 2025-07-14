package main

import (
	"fmt"
)

func insert_sort(a *[10]int, n int) {
	for pos := 1; pos < n; pos++ {
		i := pos
		for i > 0 && a[i-1] > a[i] {
			tmp := a[i]
			a[i] = a[i-1]
			a[i-1] = tmp
			i--
		}
	}
}

// сортировка вставкой
func main() {
	n := 10
	a := [10]int{7, 5, 1, 4, 2, 3, 9, 8, 0, 6}
	fmt.Println(a)
	insert_sort(&a, n)
	fmt.Println(a)
}
