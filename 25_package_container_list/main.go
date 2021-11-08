package main

import (
	"container/list"
	"fmt"
)

func main() {

	// membuat struktur data double linked list
	data := list.New()

	// menambahkan data dari belakang
	data.PushBack("Danil")
	data.PushBack("Syah")
	data.PushBack("Arihardjo")

	fmt.Println(data.Back().Value)
	fmt.Println(data.Front().Value)

	// menambahkan data dari depan
	data.PushFront("Halo")
	fmt.Println(data.Front().Value)

	fmt.Println()

	// menampilkan keseluruhan data dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println(data.Front().Next().Next().Value)

	// menampilkan data dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
