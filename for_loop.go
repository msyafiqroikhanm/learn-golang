package main

import "fmt"

// satu satunya perulangan dalam golang hanya menggnakan For Loops
func main() {

	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke-", counter)
	// }

	expression := []string{"Aja", "Banget", "Bangeettt bangeett"}
	for i := 0; i < len(expression); i++ {
		fmt.Println("Sayang Nahla", expression[i])
	}

	//ini lebih kaya foreach
	// for i, value := range expression {
	// 	fmt.Println("Index ", i, "=", value)
	// }

	//ini kalau ga butuh index (diganti jadi underscore)
	// for _, value := range expression {
	// 	// fmt.Println("Index ", i, "=", value)
	// 	fmt.Println(value)
	// }
}
