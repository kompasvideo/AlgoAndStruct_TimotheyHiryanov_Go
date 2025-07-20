package main

import "fmt"

const MAX_BINARY_TO_GENERATE = 100

var (
	digits_combination [MAX_BINARY_TO_GENERATE]int
	top                int = 0
)

func generate_binary_numbers(digit_left_to_generation int) {
	if digit_left_to_generation == 0 { // base case
		for i := 0; i < top; i++ {
			fmt.Print(digits_combination[i])
		}
		fmt.Println()
	} else { // recursive case
		digits_combination[top] = 0
		top++
		generate_binary_numbers(digit_left_to_generation - 1)
		top--

		digits_combination[top] = 1
		top++
		generate_binary_numbers((digit_left_to_generation - 1))
		top--
	}
}

// рекурсивная генерация бинарных чисел
func main() {
	fmt.Print("Enter bin number length: ")
	var n int
	fmt.Scan(&n)
	generate_binary_numbers(n)
}
