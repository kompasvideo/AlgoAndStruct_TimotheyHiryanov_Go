package main

import "fmt"

// проверка что число содержит цифру 7
func main() {
	fmt.Println("input x: ")
	var x int
	fmt.Scan(&x)
	flag := false
	for {
		flag = (x%10 == 7) || flag
		x /= 10
		if x == 0 {
			break
		}
	}
	fmt.Print("flag: ")
	fmt.Println(flag)
}
