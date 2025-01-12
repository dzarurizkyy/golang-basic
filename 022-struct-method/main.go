package main

import "fmt"

// Struct
type Customer struct {
	Name, Address string
	Age           int
}

// Method
func (customer Customer) SayHello(name string) string {
	greeting := fmt.Sprintf("Hello %s, my name is %s", name, customer.Name)
	return greeting
}

func main() {
	customer := Customer{"Dzaru", "Indonesia", 21}
	sayHello := customer.SayHello("Rizky")

	fmt.Println(sayHello)
}
