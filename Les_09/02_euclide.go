package main

import "fmt"

func euclide(a, b int) int {
	if b == 0 {
		return a
	}
	return euclide(b, a%b)
}

// рекурсия алгоритма евклида- наибольший общий делитель
// 20 6
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(euclide(a, b))
}
