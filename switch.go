package main

import (
	"fmt"
)

func main(){
	name := "rizky fa"

	switch name {
	case "rizky":
		fmt.Println("Hello rizky")
		fmt.Println("Hello rizky")
		fmt.Println("Hello rizky")
	case "wahyu":
		fmt.Println("Hello wahyu")
		fmt.Println("Hello wahyu")
		fmt.Println("Hello wahyu")
	default:
		fmt.Println("kenalan lah")
	}

	// switch length := len(name); length > 5{ //short statement
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name) // switch tanpa kondisi (kondisinya dimasukan didalam casenya)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
