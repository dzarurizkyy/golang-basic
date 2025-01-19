package main

import "fmt"

/*
	Returning Multiple Values
	- Function tidak hanya dapat mengembalikan satu value tapi juga bisa multiple value
	- Untuk memberitahu jika function mengembalikan multiple value, harus menulis semua tipe data return value nya di function
*/

/*
	Menghiraukan Return Value
	- Multiple return value wajib ditangkap semua value nya
	- Jika ingin menghiraukan return value tersebut bisa menggunakan tanda garis bawah (_)
*/

func handleName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func main() {
	fmt.Println(handleName("Dzaru", "Rizky"))
}
