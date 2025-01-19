package main

import "fmt"

/*
	Named Return Values
	- Biasanya saat membuat function yang mengembalikan value, maka hanya mendeklarasikan tipe data return value di function
	- Namun saat membuat function juga bisa membuat variable secara langsung di tipe data return function nya
*/

func handleName() (firstName, middleName, lastName string) {
	firstName = "Dzaru"
	middleName = "Rizky"
	lastName = "Fathan Fortuna"

	return firstName, middleName, lastName
}

func main() {
	fmt.Println(handleName())
}
