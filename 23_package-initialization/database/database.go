package database

import "fmt"

var connection string

func init() {
	connection = "MySQL"

	fmt.Println("init di panggil")
}

func GetDatabase() string {
	return connection
}
