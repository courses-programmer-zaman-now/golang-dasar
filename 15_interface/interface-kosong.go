package main

import "fmt"

func Ups(i int, apapun interface{}) interface{} {
	if i == 1 && apapun == "tipe-data-bebas" {
		return 1994
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func getUmur(umur interface{}) interface{} {
	return umur
}

func main() {
	var data interface{} = Ups(1, "tipe-data-bebas")
	fmt.Println(data)

	fmt.Println(getUmur(27))
	fmt.Println(getUmur("27"))
	fmt.Println(getUmur(27.2))
	fmt.Println(getUmur(true))
}
