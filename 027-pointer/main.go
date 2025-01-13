package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	address1 := Address{"Kediri", "Jawa Timur", "Indonesia"}

	/* Pointer 
	   - Kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada (pass by reference)
	*/
	
	address2 := &address1

	address2.city = "Jombang"

	fmt.Println(address1)
	fmt.Println(address2)
}
