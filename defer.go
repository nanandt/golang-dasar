package main

import "fmt"

func logging() {
	fmt.Println("Selesai memenaggil function")
}

func runApplication(value int){
	defer logging() //defer mengeskekusi func logging yg dipanggil walopun ada error di func runApplication
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(2)
}
