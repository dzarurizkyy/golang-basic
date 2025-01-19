package main

import "fmt"

/*
	Function Value
	- Function adalah first class citizen
	- Function juga merupakan tipe data dan bisa disimpan didalam variable
*/

func getGoodBye(name string) string {
	return "Good Bye, " + name + "!"
}

func main() {
	closing := getGoodBye
	fmt.Println(closing("Dzaru"))
}
