package main

import "fmt"

type hasName interface {
	getName() string
}

func sayHello(hasName hasName) {
	fmt.Println("Hello", hasName.getName())
}

type Person struct{
	Name string
}

type Animal struct {
	Name string
}

func (animal Animal) getName() string{
	return animal.Name
}

func (person Person) getName() string{
	return person.Name
}

func main() {
	rizky := Person{
		Name: "riky",
	}
	sayHello(rizky)


	cat := Animal{
		Name: "Cat",
	}
	sayHello(cat)
}
