package main

import "fmt"

func main() {
	counter := 0

	/*
		Clousure
		Kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
	*/

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
}
