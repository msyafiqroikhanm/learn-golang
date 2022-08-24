package main

import "fmt"

func main() {
	name := "Roikhan"

	if name == "Syafiq" {
		fmt.Println("Halo", name)
	} else if name == "Roikhan" {
		fmt.Println("Haii", name)
	} else {
		fmt.Println("Boleh kenalan ga?")
	}

	if length := len(name); length > 5 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("pas")
	}
}
