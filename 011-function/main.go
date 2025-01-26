package main

import "fmt"

/*
	Function
	- Sebuah blok kode yang sengaja dibuat dalam program agar bisa digunakan berulang-ulang
	- Cara membuatnya dengan menggunakan kata kunci func lalu diikuti dengan nama functionnya dan blok kode isi function nya
	- Untuk mengeksekusi function yang telah dibuat dengan memanggilnya menggunakan kata kunci nama function diikuti tanda kurung buka kurung tutup
*/

/*
	Function Parameter
	- Saat membuat function, kadang-kadang membutuhkan data dari luar atau disebut dengan parameter
	- Pada function dapat menambahkan parameter lebih dari satu
	- Parameter tidaklah wajib, jadi bisa membuat function sama seperti sebelumnya
	- Namun jika menambahkan parameter di function, maka ketika memanggil function tersebut wajib memasukkan data ke parameternya
*/

/*
	Function Return Value
	- Untuk mengetahui function mengembalikan data harus menuliskan tipe data kembalian dari function tersebut
	- Jika function telah dideklarasikan dengan tipe data pengembalian, maka wajib didalam function nya harus mengembalikan datanya
	- Untuk mengembalikan data dari function bisa menggunakan kata kunci return diikuti dengan datanya
*/

func sayHello() {
	fmt.Println("Hello!")
}

func main() {
	sayHello()
}
