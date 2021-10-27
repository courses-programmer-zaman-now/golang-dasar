package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Married       bool
}

func main() {
	// cara deklarasi ke-1
	var danil Customer
	danil.Name = "Danil"
	danil.Age = 27
	danil.Address = "Bintan"
	danil.Married = true

	fmt.Println(danil)
	if danil.Married {
		fmt.Println(danil.Name, "sudah menikah")
	} else {
		fmt.Println(danil.Name, "single")
	}

	// deklarasi slice cara ke-2
	haykal := Customer{
		Name:    "Haykal",
		Address: "Tanjung Uban",
		Age:     3,
		Married: false,
	}

	if danil.Age > haykal.Age {
		fmt.Println("danil lebih tua")
	} else {
		fmt.Println("haykal lebih tua")
	}

	// deklarasi cara ke-3 , bergantung pada urutan field slice, rentan error
	fika := Customer{"fika", "Uban", 27, true}
	fmt.Println(fika)

}
