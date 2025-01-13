package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	address1 := Address{"Kediri", "Jawa Timur", "Indonesia"}
	address2 := &address1
	*address2 = Address{"Jakarta Selatan", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}