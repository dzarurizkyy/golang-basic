package main

import "fmt"

func main() {
	/*
		Break
		Digunakan untuk menghentukan seluruh perulangan
	*/

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("Loop-%d \n", i)
	}

	fmt.Println("")

	/*
		Continue
		Digunakan untuk menghentikan perulangan yang berjalan dan langsung melanjutkan ke perulangan yang selanjutnya
	*/
	
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("Loop-%d \n", i)
	}
}
