package main

import "fmt"

// сортировка пузырьком
func main() {
	const n = 5
	a := [5]int{5, 3, 2, 1, 4}
	fmt.Println(a)
	is_sorted := false
	i := 0
	for !is_sorted {
		i = 0
		is_sorted = true
		for i < n-1 {
			if a[i] > a[i+1] {
				tmp := a[i]
				a[i] = a[i+1]
				a[i+1] = tmp
				is_sorted = false
			}
			i++
		}
	}
	fmt.Println(a)
}
