package main

import (
	"fmt"
)

func main(){
	person := map[string]string{
		"name": "rizky",
		"address": "tegal",
	}

	person["title"] = "bakul sega"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	book := make(map[string]string)
	book["title"] = "Belajar Golang Dasar"
	book["author"] = "rizky"
	book["address"] = "tegal"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book,"ups")

	fmt.Println(book)
	fmt.Println(len(book))

}
