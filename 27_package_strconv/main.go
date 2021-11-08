package main

import (
	"fmt"
	"strconv"
)

func main() {
	// konversi string - boolean
	boolean, err := strconv.ParseBool("false")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	// konversi string - int
	number, err := strconv.ParseInt("1000", 10, 16)
	if err == nil {
		fmt.Println(number + 2)
	} else {
		fmt.Println(err.Error())
	}

	v1 := "10"
	if s, err := strconv.Atoi(v1); err == nil {
		fmt.Printf("%T, %v \n", s, s)
	}

	// konversi string - float
	float, err := strconv.ParseFloat("15.5", 64)
	if err == nil {
		fmt.Println(float + 10.2)
	} else {
		fmt.Println(err.Error())
	}

	// konversi int - string
	value := strconv.FormatInt(10000, 10)
	fmt.Println(value)

	i := 10
	s := strconv.Itoa(i)
	fmt.Printf("%T, %v \n", s, s)

	// konversi float - string
	v := 3.14
	string32 := strconv.FormatFloat(v, 'E', -1, 32)
	fmt.Printf("%T, %v\n", string32, string32)

	// konversi bool - string
	strBool := strconv.FormatBool(false)
	fmt.Println(strBool)

}
