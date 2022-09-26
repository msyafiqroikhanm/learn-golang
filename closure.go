package main

import "fmt"

func main() {
	counter := 0
	name := "Syafiq"

	increment := func() {
		name := "Nahla"
		fmt.Println("increment")
		fmt.Println(name)
		counter++
	}

	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
