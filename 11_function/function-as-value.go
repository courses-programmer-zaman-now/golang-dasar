package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
     // function as value
	sayGoodBye := getGoodBye

	result := sayGoodBye("Danil")
	fmt.Println(result)
     fmt.Println(getGoodBye("haykal"))
}