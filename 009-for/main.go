package main

import "fmt"

func main() {
	names := []string{"Dzaru", "Rizky", "Fathan", "Fortuna"}

	/*
		For
		- Salah satu fitur perulangan
		- Bisa digunakan untuk melakukan iterasi terhadap data collection, contohnya array, slice, dan map
	*/

	// For (Short Statement)
	for index, name := range names {
		fmt.Printf("Index %d : %s \n", index, name)
	}
}
