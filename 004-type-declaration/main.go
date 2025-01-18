package main

import "fmt"

func main() {
	/*
		Type Declarations
		- Kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
		- Digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti
	*/

	type Password int
	var pass Password = 1234
	fmt.Println(pass)
}
