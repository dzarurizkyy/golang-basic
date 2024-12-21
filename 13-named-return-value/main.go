package main

import "fmt"

func handleName()(firstName, middleName, lastName string) {
	firstName = "Dzaru"
	middleName = "Rizky"
	lastName = "Fathan Fortuna"

	return firstName, middleName, lastName
}

func main() {
	fmt.Println(handleName())
}