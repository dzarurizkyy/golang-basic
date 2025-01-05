package main

import "fmt"

/*
	Struct
	- Sebuah template data yang digunakan untuk menggabungkan nol atau lebih
	  tipe data lainnya dalam satu kesatuan
	- Representasi data dari aplikasi yang dibuat

*/

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var kirito Customer

	kirito.Name = "Kirito"
	kirito.Address = "Aincrad"
	kirito.Age = 14

	fmt.Println(kirito)

	// Struct Literarls
	dzaru := Customer{"Dzaru", "Indonesia", 21}
	fmt.Println(dzaru)
}
