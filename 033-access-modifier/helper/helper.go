package helper

// Tidak bisa diakses dari luar
var version = "1.0.0"

// Tidak bisa diakses dari luar
func sayGoodBye(name string) string {
	return "Good Bye, " + name
}

// Bisa diakses dari luar
var Application = "Golang"

// Bisa diakses dari luar
func SayHello(name string) string {
	return "Hello, " + name
}
