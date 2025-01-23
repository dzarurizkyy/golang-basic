package main

import "fmt"

/*
	Recursive Function
	- Function yang memanggil function dirinya sendiri
	- Kadang dalam pekerjaan, untuk menyelesaikan sebuah kasus lebih mudah menggunakan recursive function dibandingkan tidak menggunakan recursive function
	- Contoh kasus yang lebih mudah diselesaikan mengggunakan recursive adalah Factorial
*/

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialRecursive(10))
}
