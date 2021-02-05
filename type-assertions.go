package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	result := random()
	// var resultString string = result.(string) //parameternya harus sesuai dengan return interface kosongnya
											// jika tidak (assigment paramnya tdk sesuai) maka akan terjadi panic
	// fmt.Println(resultString)

	switch value := result.(type){
	case string:
		fmt.Println("value",value, "is string")
	case int:
		fmt.Println("value", value, "is integer")
	default:
		fmt.Println("Unknown Type")
	}
}

/**
jika tetap merubah return interface secara paksa maka lebih baik menggunakan switch expression
agar lebih aman
*/
