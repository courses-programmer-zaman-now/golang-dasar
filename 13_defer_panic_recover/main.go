package main

import "fmt"

func endApp() {
	message := recover()
	// jika error
	if message != nil {
		fmt.Println("Error dengan message : ", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan normal")
}

func getHello(name string) string {
	return "helo " + name
}

func main() {
	runApp(true)

	println(getHello("danil syah"))
}
