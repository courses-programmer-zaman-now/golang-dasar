package main

import "fmt"

func sayHello() {
	fmt.Println("Hai danil")
}

// function parameter
func sayHelloTo(firstName string, lastName string){
     fmt.Println("hello", firstName, lastName)
}

// function return value
func getHello(name string)string{
     return "hello " + name
}

func getTambah(bil1 int8, bil2 int8) string{
     jumlah := bil1 + bil2
     if jumlah % 2 == 0 {
          return "GENAP"
     }else{
          return "GANJIL"
     }
}

// returning mutltiple values
func getGrade(firstName string, lastName string, n int)(string, int, string){
     fullName := firstName + lastName
     nilai := n
     var grade string
     if nilai >= 80 {
          grade = "A"
     }else if nilai >= 70 {
          grade = "B"
     }else if nilai >= 60 {
          grade = "C"
     }else {
          grade = "E"
     }

     return fullName, nilai, grade
}

func getFullName () (string, string, string){
     return "danil", "syah", "arihardjo"
}

// named return values
func luasSegitiga(a float32, t float32)(alas float32, tinggi float32, luas float32){
     alas = a
     tinggi = t
     luas = 0.5 * alas * tinggi
     return 
}

func main() {
     for i := 0; i <= 10; i++{
          sayHello()
     }

     sayHelloTo("danil", "syah")

     result := getHello("udinmarkudin")
     fmt.Println(result)

     keterangan := getTambah(3, 5)
     fmt.Println(keterangan)

     // returning multiple values
     fullName ,nilai, grade := getGrade("danil","syah",75)
     fmt.Println("Nama : ", fullName,"Nilai anda :", nilai, "dengan grade : ", grade)

     firstName, _, lastName := getFullName()
     fmt.Println(firstName, lastName)

     // named return values
     alas, tinggi, luas := luasSegitiga(3.4, 6.2)
     fmt.Println("alas : ", alas, "tinggi : ", tinggi, " Luas : ", luas)
}