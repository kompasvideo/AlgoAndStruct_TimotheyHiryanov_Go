package main

import "fmt"

// алгоритм - тест простоты числа
func main() {
	flag := true
	fmt.Println("input x: ")
	var x int
	fmt.Scan(&x)
	for d := 2; d < x; d++ {
		if x%d == 0 {
			flag = false
		}
	}
	fmt.Println(flag)
}
