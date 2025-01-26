package main

import "fmt"

/*
	Slice
	- Potongan dari data array
	- Mirip dengan array yang membedakan adalah ukuran array bisa berubah
	- Slice dan array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian atau seluruh data di array
*/

/*
	Detail Slice
	- Memiliki 3 tipe data, yaitu pointer, length, dan capacity
	- Pointer adalah penujuk data pertama di array
	- Length adalah panjang dari slice
	- Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity
*/

/*
	Function Slice
	- len(slice)                         : Untuk mendapatkan panjang
	- cap(slice)                         : Untuk mendapatkan kapasitas
	- append(slice, data)                : Membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
	- make([]TypeData, length, capacity) : Membuat slice baru
	- copy(destination, source)          : Menyalin slice dari source ke destination
*/

func main() {
	skills := [...]string{"HTML", "CSS", "JS"}

	slice1 := skills[1:2]
	slice2 := skills[:2]
	slice3 := skills[2:]
	slice4 := skills[:]

	fmt.Println(slice1, slice2, slice3, slice4)

	// ---------------------

	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	daySlice := days[5:]
	daySlice[0] = "New Saturday"
	daySlice[1] = "New Sunday"

	fmt.Println(days, daySlice)

	// ---------------------

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Dzaru"
	newSlice[1] = "Rizky"

	fmt.Println(newSlice, len(newSlice), cap(newSlice))

	// ---------------------

	fromSlice := newSlice[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(fromSlice, toSlice)
}
