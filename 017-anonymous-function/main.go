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
	registerUser("Dzaru", func(name string) bool { if name == "bitch" { return true }; return false })
	registerUser("bitch", func(name string) bool { if name == "bitch" { return true }; return false })
}