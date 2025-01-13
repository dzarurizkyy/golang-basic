package main

import "fmt"

type Man struct {
	Name string
}

/* Pointer di Method
	- Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
	- Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method
*/

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	customer := Man{"Dzaru Rizky"}
	customer.Married()
	fmt.Println(customer.Name)
}