package main

import "fmt"

func main() {
	// closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
	// scope variabel adalah lingkup kerja sebuah variabel
	name := "danil"
	counter := 0
	increment := func() {
		fmt.Println("Increment")
		name := "udin"
		counter++

		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
