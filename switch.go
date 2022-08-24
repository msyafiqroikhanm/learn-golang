package main

import "fmt"

func main() {
	name := "Syafiq"

	switch name {
	case "Syafiq":
		fmt.Println("Halo Syafiq")
	case "Roikhan":
		fmt.Println("Hai Roikhan")
	case "Maulana":
		fmt.Println("Helo Maulana")
	default:
		fmt.Println("Kamu siapa??")
	}

	//short statement
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("terlalu panjang")
	// case false:
	// 	fmt.Println("pas")
	// }
	//tidak ada default karena boolean hanya ada true dan false

	//switch tanpa kondisi, lebih gampang dibaca dibanding if else yang terlalu panjang
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("terlalu panjang")
	case length > 5:
		fmt.Println("lumayan panjang")
	default:
		fmt.Println("pas")
	}
}
