package main

import "fmt"

func handleName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func main() {
	fmt.Println(handleName("Dzaru", "Rizky"))
}