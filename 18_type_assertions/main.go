package main

import "fmt"

func random() interface{} {
	// return 1994
	return "danil syah"
}

func main() {
	var result interface{} = random()
	// type assertions adalah mengkonversi tipe interface menjadi tipe data string, int, float , bool
	// pastikan tipe data return sesuai dengan assertion

	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknow type")
	}

}
