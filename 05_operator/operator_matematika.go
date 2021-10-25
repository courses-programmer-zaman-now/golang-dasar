package main

import "fmt"

func main() {
	// operator aritmatika
     var tambah = 10 + 10
	var kurang = 10 - 4
	var kali = 10 * 10
	var bagi float32 = 1.0 / 2.0
	var modulus = 10 % 3
	
	fmt.Println(tambah)
	fmt.Println(kurang)
	fmt.Println(kali)
	fmt.Println(bagi)
	fmt.Println(modulus)

     // augmented assignments
     var a = 20
     
     a += 10
     fmt.Println(a)
     a -= 10
     fmt.Println(a)
     a *= 10
     fmt.Println(a)
     a /= 10
     fmt.Println(a)
     a %= 10
     fmt.Println(a)

     // unary operator
     a++
     fmt.Println(a)

     a--
     fmt.Println(a)

     var negative = -100
     var positive = +100
     fmt.Println(negative)
     fmt.Println(positive)

     var married = false
     fmt.Println(!married)

}