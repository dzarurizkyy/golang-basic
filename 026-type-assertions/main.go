package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	result := random()

	/*
		Type Assertions
		- Kemampuan mengubah tipe data menjadi tipe yang diinginkan
		- Sering digunakan ketika bertemu interface kosong
	*/

	switch result.(type) {
	case string:
		fmt.Println("String:", result)
	case int:
		fmt.Println("Int:", result)
	default:
		fmt.Println("Unknown:", result)
	}
}
