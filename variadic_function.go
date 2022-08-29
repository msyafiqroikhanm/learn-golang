package main

import "fmt"

func sumAll(numbers ...int) (result int) {
	result = 0
	for _, value := range numbers {
		result += value
	}
	return
}

func main() {
	total := sumAll(20, 20, 10)
	fmt.Println(total)

	//menggunakan slice sebagai variable arguments (varargs)
	slice := []int{28, 30, 89}
	total = sumAll(slice...)
	fmt.Println(total)
}
