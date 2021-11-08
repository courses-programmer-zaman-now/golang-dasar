package main

import "fmt"

func main() {
	// deklarasi map
	person := map[string]string{
		"name":    "danil",
		"address": "Bintan",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// menambahkan data
	person["title"] = "programmer"

	fmt.Println(person)

	// menghapus data
	delete(person, "title")
	fmt.Println(person)

	// membuat map baru
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Danil"
	book["penerbit"] = "anak baik"

	fmt.Println(book)
	fmt.Println(len(book))

	// edit data
	book["penerbit"] = "anak bangsa"
	fmt.Println(book)
}
