package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return factorial(n-1) * n
}
func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(factorial(x))
}
