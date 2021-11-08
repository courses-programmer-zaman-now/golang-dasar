package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Danil Syah Arihardjo", "Udin"))
	fmt.Println(strings.Contains("Haykal Dafiansyah", "Haykal"))

	fmt.Println(strings.Split("Danil Syah Arihardjo", " "))

	fmt.Println(strings.ToLower("Nufika Fitriani Setiawan"))
	fmt.Println(strings.ToUpper("Nufika Fitriani Setiawan"))
	fmt.Println(strings.ToTitle("danil syah"))

	kalimat := "   keep focus learn programming    "
	fmt.Println(kalimat)
	fmt.Println(strings.Trim(kalimat, " "))

	fullName := "haykal syah haykal arihardjo"
	fmt.Println(strings.ToUpper(strings.ReplaceAll(fullName, "haykal", "danil")))

}
