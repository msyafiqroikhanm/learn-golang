package main

import "fmt"

func getGoodBy(name string) string {
	return "GoodBye " + name
}

func main() {
	sayGoodBy := getGoodBy
	fmt.Println(sayGoodBy("Syafiq"))
}
