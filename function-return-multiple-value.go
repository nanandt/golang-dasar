package main

import "fmt"

func fullName() (string, string, int) {
	return "muhamad", "fatih umurnya adalah", 22
}

func main() {
	firstName, lastName, _ := fullName()
	fmt.Println(firstName, lastName)
}
