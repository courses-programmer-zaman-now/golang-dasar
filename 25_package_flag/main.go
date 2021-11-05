package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your Database host")
	var user *string = flag.String("user", "root", "Put your Database user")
	var password *string = flag.String("password", "admin", "Put your Database password")
	var number *int = flag.Int("number", 1234, "Put your number")

	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)
	fmt.Println("Number : ", *number)
}
