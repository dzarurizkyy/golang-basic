package main

import "fmt"

func sumAll(numbers ...int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

func main() {
	numbers := []int{2, 2, 2, 2, 2}

	fmt.Println(sumAll(10, 10, 10, 10, 10))
	fmt.Println(sumAll(numbers...))
}