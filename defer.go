package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

// defer berfungsi mengerjakan function ketika suatu function selesai dikerjakan, tidak peduli apakah error atau tidak
func runApplication(value int) {
	defer logging() // diletakan di atas agar dapat didefinisikan.
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(0)
}
