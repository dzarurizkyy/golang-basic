package main

import "fmt"

func main() {
	names := []string{"Dzaru", "Rizky", "Fathan", "Fortuna"}

	for index, name := range names {
		fmt.Printf("Index %d : %s \n", index, name)
	} 
}