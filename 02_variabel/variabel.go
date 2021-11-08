package main

import "fmt"

func main() {
     // penulisan deklarasi variabel
	var name string
     var age = 27
     country := "Indonesia"

     // penulisan deklarasi multiple variabel
     var (
          firstName = "Danil"
          lastName = "Syah"
     )

     fmt.Println(firstName +  lastName)

	name = "Danil syah arihardjo"
	fmt.Println("name : ", name)

     name = "haykal"
     fmt.Println(name)
     fmt.Println(age)
     fmt.Println(country)
}