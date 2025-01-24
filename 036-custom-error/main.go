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
	err := SaveData("", nil)

	if err != nil {
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println(finalError.Error())
		case *notFoundError:
			fmt.Println(finalError.Error())
		default:
			fmt.Println("Unknown Error")
		}
	} else {
		fmt.Println("Success")
	}
}
