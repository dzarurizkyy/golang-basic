package main

import (
	"errors"
	"fmt"
)

func Division(value int, division int) (int, error) {
	if division == 0 {
		return 0, errors.New("division by zero")
	} else {
		return value / division, nil
	}
}

func main() {
	result, err := Division(100, 0)

	if err == nil {
		fmt.Println("Result:", result)
	} else {
		fmt.Println("Error:", err.Error())
	}
}