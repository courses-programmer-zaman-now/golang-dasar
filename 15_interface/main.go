package main

import "fmt"

type HasName interface {
	GetName() string
	GetAge() int
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func ViewAge(hasName HasName) {
	fmt.Println("umur ku sekarang", hasName.GetAge())
}

type Person struct {
	Name string
	Age  int
}

func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetAge() int {
	return person.Age
}

type Animal struct {
	Name string
	Age  int
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (animal Animal) GetAge() int {
	return animal.Age
}

func main() {
	var person1 Person
	person1.Name = "Danil Syah"
	person1.Age = 27
	SayHello(person1)
	ViewAge(person1)

	animal1 := Animal{Name: "Kucing"}
	SayHello(animal1)
}
