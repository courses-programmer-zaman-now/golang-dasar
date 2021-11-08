package main

import "fmt"

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
func main() {
	// function recursive adalah , pemanggilan function dirinya sendiri di dalam function dirinya
	result := factorialRecursive(5)
	fmt.Println(result)
}
