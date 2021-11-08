package main

import "fmt"

func main() {
	name := "Haykal"

	switch name {
	case "Haykal":
		fmt.Println("Hello Haykal")
	case "Joko":
		fmt.Println("hai fika")
	default:
		fmt.Println("nama tidak ada")
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi
	nilai := 78
	switch {
	case nilai >= 80:
		fmt.Println("A")
	case nilai >= 75:
		fmt.Println("B")
	case nilai >= 60:
		fmt.Println("C")
	default:
		fmt.Println("Perbaiki nilai anda")
	}
}
