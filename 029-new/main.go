package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	address1 := new(Address)
	address2 := address1

	address2.country = "Indonesia"

	fmt.Println(address1)
	fmt.Println(address2)
}
