package main

import "fmt"

//buat alias untuk detail function
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Syafiq", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
