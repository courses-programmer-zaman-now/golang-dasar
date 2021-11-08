package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("d([a-z])n")

	fmt.Println(regex.MatchString("danil"))
	fmt.Println(regex.MatchString("dUnil"))

	search := regex.FindAllString("dai den don dan", 1)
	fmt.Println(search)
}
