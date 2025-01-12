package main

import "fmt"

type Filter func(string) string

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
