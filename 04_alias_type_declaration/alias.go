package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var noKtp NoKtp = "3212902929221"
	var status Married = true

	fmt.Println(noKtp)
	fmt.Println(status)
}