package main

import "fmt"

func main() {
	bilangan := -2

	if bilangan > 0 {
		if bilangan%2 == 0 {
			fmt.Println("Genap")
		} else {
			fmt.Println("Ganjil")
		}
	} else {
		fmt.Println("Angka tidak boleh bilangan negatif atau Nol")
	}

	name := "test"

	if name == "danil" {
		fmt.Println("hai", name)
	} else if name == "udin" {
		fmt.Println("kenalan dong")
	} else {
		fmt.Println("tidak kenal")
	}

	// short statement if
	if length := len(name); length < 3 {
		fmt.Println("nama anda pendek")
	} else {
		fmt.Println("nama anda panjang")
	}
}
