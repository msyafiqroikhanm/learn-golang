package main

import "fmt"

/*
struck merupakan template data yang digunakan untuk menggabungkan nol atau lebih tipe data ke dalam satu kesatuan
kalau di oop ada class, struct mirip class.
*/

type Customer struct {
	Name, Address string
	Age           int
	// married       bool
}

//function struct yang seakan akan milik struct, padahal function tersebut berdiri sendiri
func (customer Customer) sayHelloStruct(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var syafiq Customer
	syafiq.Name = "Syafiq"
	syafiq.Address = "Rawamangun"
	syafiq.Age = 21

	nahla := Customer{
		Name:    "Nahla",
		Address: "Dukuhdamu",
		Age:     22,
	}

	syean := Customer{"Syean", "Babakan", 9}

	fmt.Println(syafiq)
	fmt.Println(nahla)
	fmt.Println(syean)

	syafiq.sayHelloStruct("Nahla") //pemanggilan struct method
}
