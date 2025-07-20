package main

import "fmt"

// ханойские башни
func main() {
	hanoi(1, 2, 4)
}

/**
* Hanoi solution finder
* param i: start pin number
* param k: finish pin number
* param n: number of disk
 */
func hanoi(i, k, n int) {
	if n == 1 {
		fmt.Printf("Move disk 1 from pin %d to pin %d \n", i, k)
	} else {
		tmp := 6 - i - k
		hanoi(i, tmp, n-1)
		fmt.Printf("Move disk %d from pin %d to pin %d \n", n, i, k)
		hanoi(tmp, i, n-1)
	}
}
