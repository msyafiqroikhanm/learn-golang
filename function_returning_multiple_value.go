package main

import "fmt"

func getFullName() (string, string, string) {
	return "Syafiq", "Roikhan", "Maulana"
}

func main() {
	firstname, _, lastname := getFullName()
	// untuk menghiraukan salah satu variable bisa menggunakan underscored _
	fmt.Println(firstname)
	fmt.Println(lastname)
}
