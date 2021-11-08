package main

import "fmt"

type Filter func(string, string) string

func sayHelloWithFilter(name string, genre string, filter Filter) {
	nameFiltered := filter(name, genre)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string, genre string)string{
    if name == "anjing" {
         if genre == "female" {
              return "..."
         }else{
              return "cowo biasa nyebut anjing"
         }
    }else{
         return name
    }
}

func main() {
     sayHelloWithFilter("anjing","female", spamFilter)
     sayHelloWithFilter("anjing", "male", spamFilter)
}