package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye, " + name + "!"
}

func main() {
	closing := getGoodBye
	fmt.Println(closing("Dzaru"))
}