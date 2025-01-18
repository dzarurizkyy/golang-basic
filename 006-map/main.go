package main

import "fmt"

func main() {
	/*
		Map
		- Berisikan kumpulan data yang sama, namun bisa menentukan jenis tipe data index yang digunakan
		- Tipe data kumpulan key-value, dimana key bersifat unik, tidak boleh sama
		- Berbeda dengan array dan slice, jumlah data yang dimasukkan ke dalam map boleh sebanyak-banyaknya, 
		  asalkan key berbeda, jika menggunakan key sama, maka secara otomatis data sebelumnya akan diganti 
		  dengan data yang baru
	*/

	/*
		Function Map
		- len(map)                     : Untuk mendapatkan jumlah data di map
		- map[key]                     : Mengambil data di map dengan key
		- map[key] = value             : Mengubah data di map dengan key
		- make(map[TypeKey]TypeValue)  : Membuat map baru
		- delete(map, key)             : Menghapus data di map dengan key
	*/
	
	account := map[string]string{
		"name": "Dzaru Rizky Fathan Fortuna",
		"role": "Junior Web Developer",
		"exp":  "1 years",
	}

	fmt.Println(account)

	delete(account, "exp")

	fmt.Println(account)
}
