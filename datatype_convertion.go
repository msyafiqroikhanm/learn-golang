package main

import "fmt"

func main() {
	var nilai32 int32 = 3245
	var nilai16 int16 = int16(nilai32)
	var nilai64 int64 = int64(nilai16)
	var nilai8 int8 = int8(nilai32)
	var nilaiU8 uint8 = uint8(nilai32)

	fmt.Println(nilaiU8)
	fmt.Println(nilai8)
	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai64)

	var name = "Syafiq"
	var e = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
