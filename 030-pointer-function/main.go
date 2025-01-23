package main

import "fmt"

type Address struct {
	city, province, country string
}

/*
	Pointer di Function
	- Saat membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
	- Oleh karena itu, jika mengubah data di dalam function, data aslinya tidak akan berubah
	- Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
	- Namun bila ingin membuat function yang bisa mengubah data asli parameter tersebut, dapat menggunakan pointer di function
	- Untuk menjadikan sebuah parameter sebagai pointer menggunakan operator * di parameternya
*/

func ChangeCountryToIndonesia(address *Address) {
	address.country = "Indonesia"
}

func main() {
	address1 := Address{}
	ChangeCountryToIndonesia(&address1)
	fmt.Println(address1)
}