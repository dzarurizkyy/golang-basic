package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	address1 := Address{"Kediri", "Jawa Timur", "Indonesia"}
	address2 := &address1

	/*
		Asterisk Operator
		- Saat mengubah variable pointer, maka yang berubah hanya variable tersebut
		- Semua variable yang mengacu ke data yang sama tidak akan berubah
		- Jika ingin mengubah seluruh variable yang mengacu ke data tersebut, data menggunakan asterisk operator (*)
	*/

	*address2 = Address{"Jakarta Selatan", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
