package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

// Custom Error
func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "dzaru" {
		return &notFoundError{"data not found"}
	}

	return nil
}

func main() {
	err := SaveData("dzaru", nil)

	if err != nil {
		if validationError, ok := err.(*validationError); ok {
			fmt.Println(validationError.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println(notFoundError.Error())
		}
	} else {
		fmt.Println("Success")
	}
}
