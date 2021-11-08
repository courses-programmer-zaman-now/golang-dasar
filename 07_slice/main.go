package main

import "fmt"

func main() {
	// slice adalah potongan dari array
	// jumlah data dalam slice bisa berubah-ubah

	// deklarasi array
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

	// jika data array di ubah maka data pada slice pun ikut terubah
	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// begitupun pada data slice yang mengacu ke array , jika ada data di ubah maka data pada array
	// akan ter update juga
	// slice1[0] = "mei-diubah"
	// fmt.Println(months)

	var slice2 = months[11:]
	fmt.Println(slice2)

	// menambahkan data diakhir element, jika capasitas nya sudah penuh maka, akan dibuat array baru
	var slice3 = append(slice2, "Danil")
	fmt.Println(slice3)
	fmt.Println(months)
	slice3[1] = "Bukan-Desember"
	fmt.Println(slice3)

	// slice2[1] = "test"
	fmt.Println(slice2)
	fmt.Println(months)

	// membuat slice baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Danil"
	newSlice = append(newSlice, "syah")

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	var sliceNew1 = append(newSlice, "ari")
	fmt.Println(sliceNew1)
	sliceNew1[2] = "coba"
	fmt.Println(newSlice)
	fmt.Println(sliceNew1)

	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// hati-hati dalam membuat array dan slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
