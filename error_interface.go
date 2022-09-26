package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	// var contohError = errors.New("Upsss")
	// fmt.Println(contohError.Error())
	hasil, err := Pembagi(100, 10)
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
