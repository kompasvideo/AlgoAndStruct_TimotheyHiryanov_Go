package main

import "fmt"

// дурацкая сортировка
func main() {
	const n = 5
	a := [5]int{5, 3, 2, 1, 4}
	fmt.Println(a)
	for i := 0; i < n-1; {
		if a[i] > a[i+1] {
			tmp := a[i]
			a[i] = a[i+1]
			a[i+1] = tmp
			i = 0
		} else {
			i++
		}
	}
	fmt.Println(a)
}
