package main

import "fmt"

func endApp() {
	fmt.Println("End App...")

	message := recover(); 

	fmt.Println(message)
	fmt.Println("Restart Server...")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Internal Server Error (500)")
	}

}

func main() {
	runApp(true)
	fmt.Println("Dzaru Rizky Fathan Fortuna")
}
