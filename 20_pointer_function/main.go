package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}
func main() {
	var alamat Address = Address{
		City:     "Bandung",
		Province: "Jawa Barat",
		Country:  "",
	}

     var alamatPointer *Address = &alamat

	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)
}