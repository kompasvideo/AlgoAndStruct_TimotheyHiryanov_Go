package main

import "fmt"

func sort_radix(a *[10]int, n int) {
	var a0 [10]int
	var a1 [10]int
	for radix := 0; radix < 32; radix++ {
		a0_size := 0
		a1_size := 0
		for i := 0; i < n; i++ {
			if (a[i] & (1 << radix)) == 0 {
				a0[a0_size] = a[i]
				a0_size++
			} else {
				a1[a1_size] = a[i]
				a1_size++
			}
		}
		for i := 0; i < a0_size; i++ {
			a[i] = a0[i]
		}
		for i := 0; i < a1_size; i++ {
			a[i+a0_size] = a1[i]
		}
	}
}

// поразрядная сортировка
func main() {
	const n = 10
	a := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(a)
	sort_radix(&a, n)
	fmt.Println(a)
}
