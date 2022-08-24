package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// perubahan data di array utama akan merubah daya yang ada di slice karena terkoneksi
	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// //begitupun ketika data di slice diubah maka akan merubah data yang ada di array utama karena terkoneksi
	// slice1[0] = "Diubah Lagi"
	// fmt.Println(months)

	// var slice2 = months[2:4]
	var slice2 = months[10:]
	fmt.Println(slice2)

	//ketika melakukan append dalam keadaan masih tersedia kapasitas untuk menyimpan array tersebut maka akan mengganti dari array yang tersedia
	//tetapi ketika melakukan append dalam keadaan kapasitasnya sudah penuh maka akan membuat array baru yang tidak terkoneksi dengan array utama sebelumnya.
	var slice3 = append(slice2, "Syafiq")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Syafiq"
	newSlice[1] = "Roikhan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//pastikan panjang sama dengan array yang akan dicopy
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//ARRAY VS SLICE
	iniArray := [...]int{1, 2, 3, 4, 5} //didefinisikan len nya
	iniSlice := []int{1, 2, 3, 4, 5}    //tidak didefinisikaan len nya
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
