package main

import "fmt"

/*
	Variadic Function
	- Parameter yang berada di posisi terakhir memiliki dijadikan sebuah varargs
	- Varargs artinya datanya bisa menerima lebihd dari satu input, sama seperti array
*/

/*
	Perbedaan Array dan Varargs
	- Jika parameter tipe array wajib membuat array terlebih dahulu sebelum mengirimkan ke function
	- Jika parameter menggunakan varargs bisa langsung mengirim data nya, jika lebih dari satu cukup gunakan tanda koma
*/

func sumAll(numbers ...int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

func main() {
	numbers := []int{2, 2, 2, 2, 2}

	fmt.Println(sumAll(10, 10, 10, 10, 10))
	fmt.Println(sumAll(numbers...))
}
