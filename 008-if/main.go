package main

import "fmt"

func main() {
	const name string = "Dzaru Rizky Fathan Fortuna"

	/*
		If Expression
		- Salah satu kata kunci yang digunakan untuk percabangan
		- Percabanyan artinya bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
	*/

	/*
		Else Expression
		- Blok if akan dieksekusi ketika kondisi if bernilai true
		- Else expression akan dieksekusi oleh program jika kondisi if bernilai false
	*/

	// If (Short Statement)
	if lenght := len(name); lenght < 5 {
		fmt.Println("Name is too short")
	} else {
		fmt.Println("Name is too long")
	}
}
