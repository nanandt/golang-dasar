package main

import "fmt"

// Filter is ...
type Filter func(string) string

// func sayHelloWithName(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }

func sayHelloWithName(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string)string{
	if name == "anjing"{
		return "..."
	}

	return name
}


func main(){
	sayHelloWithName("Rizky", spamFilter)
	sayHelloWithName("anjing", spamFilter)


}
