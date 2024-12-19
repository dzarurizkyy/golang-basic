package main

import "fmt"

func main() {
	account := map[string]string{
		"name": "Dzaru Rizky Fathan Fortuna",
		"role": "Junior Web Developer",
		"exp":  "1 years",
	}

	fmt.Println(account)

	delete(account, "exp")

	fmt.Println(account)
}
