package main

import "fmt"

func main() {

	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
          counter++
	}

     // for dengan statement
     for i := 1; i <= 10; i++{
          fmt.Println("I like coding")
     }

     // for range
     slice := []string{"danil","haykal","fika","alin"}

     for i, value := range slice{
          fmt.Println("Index ", i, " = ", value)
     }

     person := make(map[string]string)
     person["nama"] = "danil syah"
     person["title"] = "programmer"
     person["address"] = "Bintan"

     for key, value := range person{
          fmt.Println(key," : ", value)
     }
}