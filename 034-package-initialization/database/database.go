package database

var connection string

/*
	Package Initialization
	- Saat membuat package dapat membuat sebuah function yang akan diakses ketika package tersebut diakses
	- Ini sangat cocok ketika contohnya jika package berisi function-function untuk berkomunikasi dengan database
	- Untuk membuat function yang diakses secara otomatis ketika package diakses cukup membuat function dengan nama init
*/

func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
