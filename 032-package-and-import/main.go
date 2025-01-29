package main

/*
	Import
	- Secara standar, file Golang hanya bisa mengakses file Golang lainnya yang berada dalam package yang sama
	- Jika ingin mengakses file Golang yang berada di luar package, maka dapat menggunakan import
*/

import (
	"fmt"
	"package-and-import/helper"
)

func main() {
	result := helper.SayHello("Dzaru")
	fmt.Println(result)
}
