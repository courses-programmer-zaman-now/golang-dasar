package main

import "fmt"

func main() {
	// && and
	// || or
	// ! kebalikan atau negasi

	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi
	var lulus2 bool = lulusNilaiAkhir || lulusAbsensi

	fmt.Println(lulus)
	fmt.Println(lulus2)
	fmt.Println(nilaiAkhir >= 80 && absensi >= 75)

}
