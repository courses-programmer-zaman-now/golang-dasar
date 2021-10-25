package main

import "fmt"

func main() {
	var values = [3]int{
		90,
		95,
		88,
	}

	fmt.Println(values)
     fmt.Println(values[0])
     fmt.Println(values[1])
     fmt.Println(values[2])

     // mengubah data
     values[2] = 55
     fmt.Println(values)

     var names = [3]string{
          "danil",
          "haykal",
          "fika",
     }

     fmt.Println(names)
     
     var panjangNames = len(names)
     fmt.Println(panjangNames)

     var panjangValues = len(values)
     fmt.Println(panjangValues)
}