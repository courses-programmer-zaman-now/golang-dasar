package main

import (
	"fmt"
	"time"
)

func main() {
	// mendapatkan waktu saat ini
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	// mendapatkan waktu secara manual
	timeLocal := time.Date(2021, 11, 5, 10, 53, 24, 12, time.Now().Location())
	fmt.Println(timeLocal)

	// parsing string - date
	layout := "2006-01-02" //format
	parse, _ := time.Parse(layout, "1994-10-15")
	fmt.Println(parse.Month())
}
