package main

import "fmt"

func endApp() {
	fmt.Println("End App...")

	/*
		Recover
		- Function yang digunakan untuk menangkap data panic
		- Dengan recover, proses panic akan terhenti, sehingga program akan tetap berjalan
	*/

	message := recover()

	fmt.Println(message)
	fmt.Println("Restart Server...")
}

func runApp(error bool) {
	/*
		Defer
		- Function yang dijadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi
		- Defer akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi
	*/

	defer endApp()

	/*
		Panic
		- Function yang digunakan untuk menghentikan program
		- Function yang dipanggil ketika terjadi panic pada saat program berjalan
		- Saat panic function dipanggil, program akan terhenti, namun defer function akan tetap dieksekusi
	*/

	if error {
		panic("Internal Server Error (500)")
	}
}

func main() {
	runApp(true)
	fmt.Println("Dzaru Rizky Fathan Fortuna")
}
