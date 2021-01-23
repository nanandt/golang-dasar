package main

import (
	"fmt"
)

func main(){
	var name [3]string

	name[0] = "Muhamad"
	name[1] = "Fatih"
	name[2] = "Rizky"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int{
		20,
		30,
		50,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(name))
	fmt.Println(len(values))

	var lagi [20]string
	fmt.Println(len(lagi))
}
