package main

import "fmt"

func choice_sort(a *[10]int, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if a[j] < a[i] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp
			}
		}
	}
}

// сортировка выбором
func main() {
	n := 10
	a := [10]int{9, 4, 1, 2, 8, 7, 0, 5, 3, 6}
	fmt.Println(a)
	choice_sort(&a, n)
	fmt.Println(a)
}
