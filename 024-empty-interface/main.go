package main

import "fmt"

/*
	Interface Kosong (interface{})
	- Interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya
	- Memiliki type alias bernama any
*/

func Ups() any {
	return "OK"
}

func main() {
	msg := Ups()
	fmt.Println(msg)
}
