package main

import (
	"fmt"
	"math"
)

func main() {
	// membulatkan ke nilai terdekat
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.4))

	// membulatkan nilai ke bawah
	fmt.Println(math.Floor(1.8))

	// membulatkan nilai ke atas
	fmt.Println(math.Ceil(1.3))

	var values1 float64 = 40
	var values2 float64 = 22
	nilaiTerbesar := math.Max(values1, values2)
	fmt.Println(nilaiTerbesar)

	nilaiTerkecil := math.Min(values1, values2)
	fmt.Println(nilaiTerkecil)

}
