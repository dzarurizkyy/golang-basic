package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {

	/*
		New
		- Sebelumnya untuk membuat pointer dengan mengggunakan operator &
		- Golang juga memiliki function new yang bisa digunakan untuk membuat pointer
		- Namun function new hanya mengembalikan pointer ke data kosong, artinya data tidak ada di tahap awal
	*/

	address1 := new(Address)
	address2 := address1

	address2.country = "Indonesia"

	fmt.Println(address1)
	fmt.Println(address2)
}
