package main

import "fmt"

func main() {
	var name string

	var (
		firstName = "M. Syafiq"
		lastName  = "Roikhan Maulana"
		IPK       = 3.72
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(IPK)

	name = "Syafiq Roikhan"
	fmt.Println(name)

	name = "Syafiq Maulana"
	fmt.Println(name)

	fullName := "M. Syafiq Roikhan Maulana"
	fmt.Println(fullName)

	var girlFriend = "Nahla"
	fmt.Println(girlFriend)

	var age = 21
	fmt.Println(age)
}
