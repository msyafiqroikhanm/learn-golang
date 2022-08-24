package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Syafiq"
	names[1] = "Roikhan"
	names[2] = "Maulana"

	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int8{90, 95, 80}
	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
}
