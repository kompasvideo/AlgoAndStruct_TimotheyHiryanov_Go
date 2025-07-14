package main

import "fmt"

// разложить число на множители
func main() {
	var x int
	_, _ = fmt.Scan(&x)
	var d = 2
	for x != 1 {
		if x%d == 0 {
			fmt.Println(d)
			x = x / d
		} else {
			d++
		}
	}
}
