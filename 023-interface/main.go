package main

import "fmt"

/*
	Interface
	- Tipe data Abstract, tidak memiliki implementasi langsung
 	- Berisikan definisi-definisi method
	- Digunakan sebagai kontrak
	- Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis diaggap interface tersebut
*/

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello,", hasName.GetName())
}

// Struct
type Person struct {
	Name string
}

// Method
func (person Person) GetName() string {
	return person.Name
}

func main() {
	person := Person{Name: "Dzaru"}
	sayHello(person)
}
