package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Nama   string `required:"true" max:"5"`
	Alamat string `required:"true" max:"100"`
	Umur   int    `required:"true" max:"2"`
}

// kode validation library
func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Danil", "Bintan", 27}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)

	// mendapatkan data field struct ke 0
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(1).Name)
	fmt.Println(sampleType.Field(2).Name)
	// mendapatkan jumlah field struct
	fmt.Println(sampleType.NumField())

	// mendapatkan data tag field
	fmt.Println(structField.Tag.Get("required"))
	fmt.Println(structField.Tag.Get("max"))

	// uji validasi
	fmt.Println(IsValid(sample))

	sample.Nama = ""
	validation := IsValid(sample)
	fmt.Println(IsValid(sample))

	if validation {
		fmt.Println("data sudah valid")
	} else {
		fmt.Println("data tidak boleh ada kosong")
	}

}
