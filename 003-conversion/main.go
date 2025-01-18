package main

import "fmt"

func main() {
	const name = "Dzaru"

	firstChar := name[0]
	fmt.Println(firstChar)

	firstCharStr := string(firstChar)
	fmt.Println(firstCharStr)
}
