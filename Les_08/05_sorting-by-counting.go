package main

import (
	"fmt"
	"math/rand"
)

func sort(a *[10]int, n int) {
	var f [10]int
	for i := 0; i < n; i++ {
		f[a[i]]++
	}
	i := 0
	for x := 0; x < 10; x++ {
		for k := 0; k < f[x]; k++ {
			a[i] = x
			i++
		}
	}
}
func generate_random_array(a *[10]int, n, m int) {
	for i := 0; i < n; i++ {
		a[i] = rand.Int() % m
	}
}

// сортировка подсчётом
func main() {
	n := 10
	var a [10]int
	generate_random_array(&a, n, 10)
	fmt.Println(a)
	sort(&a, n)
	fmt.Println(a)
}
