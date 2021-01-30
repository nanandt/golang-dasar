package main

import "fmt"

func getHello(name string, age int) string {
	hasil := "Rizky"

	fmt.Println(name,age)


	return hasil
}


func main() {
	result := getHello("Rizky",22)

	fmt.Println(result)
}
