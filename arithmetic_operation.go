package main

import "fmt"

func main() {
	var a = 15
	var b = 23
	var c = a + b
	var d = a * b

	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(d)

	//augmented assigment
	a += b
	d *= a

	fmt.Println(a)
	fmt.Println(d)

	var i = 1
	i++
	fmt.Println(i)

	var positive = 100
	var negative = -100
	fmt.Println(positive)
	fmt.Println(negative)
}
