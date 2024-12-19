package main

import "fmt"

func main() {
	type Password int
	var pass Password = 1234
	fmt.Println(pass)
}