package main

import "fmt"

func main() {
	/*
		Array
		- Tipe data yang berisikan kumpulan data dengan tipe data yang sama
		- Saat membuat array perlu menentukan jumlah data yang bisa ditampung oleh array tersebut
		- Daya tampung array tidak bisa bertambah setelah array dibuat
	*/

	number := [...]int{1, 2, 3, 4}
	fmt.Println(number)
}
