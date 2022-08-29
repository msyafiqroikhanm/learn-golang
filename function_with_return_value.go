package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Ngab"
	} else {
		return "Hello " + name
	}
}

func main() {
	name := getHello("Syafiq")
	fmt.Println(name)
	name = getHello("")
	fmt.Println(name)
}
