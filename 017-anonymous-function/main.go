package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Your are blocked") 
	} else {
		fmt.Println("Welcome,", name, "!")
	}
}

func main() {
	/*
		Anonymous Function
		- Sebelumnya setiap membuat function selalu memberikan sebuah nama pada function tersebut
		- Namun ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membuat function terlebih dahulu
		- Hal tersebut dinamakan anonymous function	atau function tanpa nama
	*/

	registerUser("Dzaru", func(name string) bool { if name == "bitch" { return true }; return false })
	registerUser("bitch", func(name string) bool { if name == "bitch" { return true }; return false })
}