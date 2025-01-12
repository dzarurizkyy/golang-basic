package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("Loop-%d \n", i)
	}

	fmt.Println("")

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Printf("Loop-%d \n", i)
	}
}