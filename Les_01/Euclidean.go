package main

import "fmt"

func euclid_gcd(a, b int) int {
	for {
		if a > b {
			a -= b
		} else {
			b -= a
		}
		if a == b {
			break
		}
	}
	return a
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Print("GCD(x,y) = ")
	res := euclid_gcd(x, y)
	fmt.Println(res)
}
