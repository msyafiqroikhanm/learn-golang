package main

import "fmt"

func sayHelloTo(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}

func main() {
	sayHelloTo("Syafiq", "Roikhan")
	sayHelloTo("Nahla", "Isna")
}
