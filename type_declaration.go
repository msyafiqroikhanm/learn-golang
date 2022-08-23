package main

import "fmt"

func main() {
	type noKTP string
	type married bool

	var noKTPSyafiq noKTP = "31213231231231"
	var status married = false
	fmt.Println(noKTPSyafiq)
	fmt.Println(status)
}
