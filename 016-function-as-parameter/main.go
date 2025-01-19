package main

import "fmt"

type Filter func(string) string

/*
	Function as Parameter
	- Function tidak hanya disimpan didalam variable sebagai value
	- Namun juga bisa digunakan sebagai parameter untuk function lain
*/

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Printf("Hello, %s\n", filteredName)
}

func spamFilter(name string) string {
	if name == "bitch" {
		return "..."
	}
	return name + "!"
}

func main() {
	sayHelloWithFilter("Dzaru", spamFilter)
	sayHelloWithFilter("bitch", spamFilter)
}
