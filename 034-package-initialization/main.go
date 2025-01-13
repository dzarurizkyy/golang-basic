package main

import (
	"fmt"
	"package-initialization/database"

	/*
		Blank Identifier
		- Ketika hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function yang ada di package
		- Secara default, Golang akan komplen ketika ada package yang diimport namun tidak digunakan
		- Untuk menagani hal tersebut dapat menggunakan blank identifier (_) sebelum nama package ketika melakukan import
	*/

	_ "package-initialization/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
