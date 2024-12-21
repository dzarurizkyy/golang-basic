package main

import "fmt"

func main() {
	skills := [...]string{"HTML", "CSS", "JS"}

	slice1 := skills[1:2]
	slice2 := skills[:2]
	slice3 := skills[2:]
	slice4 := skills[:]

	fmt.Println(slice1, slice2, slice3, slice4)
	
	// ---------------------

	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	daySlice := days[5:]
	daySlice[0] = "New Saturday"
	daySlice[1] = "New Sunday"

	fmt.Println(days, daySlice)

	// ---------------------

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Dzaru"
	newSlice[1] = "Rizky"

	fmt.Println(newSlice, len(newSlice), cap(newSlice))

	// ---------------------

	fromSlice := newSlice[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(fromSlice, toSlice)
}