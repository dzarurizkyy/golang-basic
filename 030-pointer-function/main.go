package main

import "fmt"

type Address struct {
	city, province, country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.country = "Indonesia"
}

func main() {
	address1 := Address{}
	ChangeCountryToIndonesia(&address1)
	fmt.Println(address1)
}