package main

import "fmt"

/*
	Nil
	- Data kosong
	- Hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer, dan channel
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": "Dzaru Rizky Fathan Fortuna",
		}
	}
}

func main() {
	data := NewMap("")

	if data == nil {
		fmt.Println("Empty Data")
	} else {
		fmt.Println(data["name"])
	}
}
