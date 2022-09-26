package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Message", message)
	}
	fmt.Println("Aplikasi selesai")
}

// panic digunakan untuk menjalankan sesuatu ketika aplikasi error/panic
// function recover yang bisa digunakan untuk menangkap data panic
// dengan adanya recover, proses panic akan terhenti sehingga program akan tetap berjalan
func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Next ...")
}
