package main

import (
	// _ "test/database"  // black identifier : hanya memanggil func init saja
	"fmt"
	"test/database"
	"test/helper"
)

func main() {
	helper.SayHello()
	helper.Perkenalan("haykal")

	var getDatabase = database.GetDatabase()
	fmt.Println(getDatabase)
}
