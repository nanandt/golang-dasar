package main

import (
	"fmt"
)

func main(){
	var name1 = "rizky"
	var name2 = "rizky"

	var result bool = name1 == name2
	fmt.Println(result)

	name1 = "wahyu"
	var n byte = name1[0]
	var nString string = string(n)
	fmt.Println(nString)

	name2 = "hadi"
	var h byte = name2[0]
	var hString string = string(h)
	fmt.Println(hString)

	var hasil bool = nString == hString
	fmt.Println(hasil)

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
