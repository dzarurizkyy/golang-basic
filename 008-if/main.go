package main

import "fmt"

func main() {
	const name string = "Dzaru Rizky Fathan Fortuna"

	if lenght := len(name); lenght < 5 {
		fmt.Println("Name is too short")
	} else {
		fmt.Println("Name is too long")
	}
}
