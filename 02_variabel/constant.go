package main

import "fmt"

func main() {
	// constant adalah variabel yang nilainya tidak bisa diubah setelah diberi nilai

	// deklarasi multiple constant
	const (
		firstName = "Danil"
		lastName  = "Syah"
		value     = 1000
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
