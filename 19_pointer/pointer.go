package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
     // pass by value , hanya menduplicate data
	address2 := address1
     
     // menggunakan pointer, pass by reference
     address3 := &address1

     // pass by value
     address2.City = "Bandung"

     // pointer membuat variable pass by reference karena langsung mengakses ke memory
     address3.City = "Sumedang"

     *address3 = Address{"Wado", "Jawa Barat", "Indonesia"}


	fmt.Println(address1)
     fmt.Println(address2)
     fmt.Println(address3)

     var address4 *Address = new(Address)
     address4.City = "Jogjakarta"
     address4.Country = ""
     fmt.Println(address4)

}