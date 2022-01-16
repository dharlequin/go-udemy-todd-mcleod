package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(numbers...))
	fmt.Println(bar(numbers))
}

func foo(numbers ...int) int {
	return bar(numbers)
}

func bar(numbers []int) int {
	var total = 0
	for _, number := range numbers {
		total += number
	}

	return total
}
