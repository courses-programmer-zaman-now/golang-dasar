package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {

	// anonymous function
	blacklistUser := func(name string) bool {
		return name == "root"
	}

	registerUser("root", blacklistUser)
	registerUser("danil", blacklistUser)
}
