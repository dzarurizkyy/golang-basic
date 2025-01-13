package main

import (
	"access-and-modifier/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Dzaru")

	fmt.Println(result)
	fmt.Println(helper.Application)
}