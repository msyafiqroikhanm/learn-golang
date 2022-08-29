package main

import "fmt"

func getCompleteName() (firstname string, middlename string, lastname string) {
	firstname = "Syafiq"
	middlename = "Roikhan"
	lastname = "Maulana"

	return
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
