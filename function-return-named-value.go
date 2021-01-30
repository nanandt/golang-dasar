package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "muhamad"
	middleName = "fatikhur"
	lastName = "rizky"

	return
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)
}
