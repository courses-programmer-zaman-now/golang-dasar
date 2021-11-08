package main

import "fmt"

// function variadic yaitu parameter function yang berupa array, jadi kita bisa di masukan dinamis
func sumAll(numbers ...int)int{
     total := 0
     for _, value := range numbers {
          total += value
     }
     return total
}

func main() {
     total := sumAll(40,53,2,6)
     fmt.Println(total)

     // slice sebagai parameter
     slice := []int{3,6,2,1,8,5}
     total = sumAll(slice...)
     fmt.Println(total)
}