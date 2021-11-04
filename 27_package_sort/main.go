package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Danil", 27},
		{"Fika", 26},
		{"Haykal", 3},
		{"Icih", 60},
	}

	// data sebelum di urutkan
	fmt.Println(users)

	// mengurutkan data dari kecil ke besar
	sort.Sort(UserSlice(users))

	// data setelah di urutkan
	fmt.Println(users)
}
