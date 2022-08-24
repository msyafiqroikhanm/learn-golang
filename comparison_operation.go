package main

import "fmt"

func main() {
	var name1 = "Syafiq"
	var name2 = "Syafiq"
	var name3 = "Nahla"

	var result bool = name1 == name2
	fmt.Println(result)

	result = name1 == name3
	fmt.Println(result)

	result = name1 != name3
	fmt.Println(result)
}
