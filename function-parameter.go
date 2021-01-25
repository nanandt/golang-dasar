package main

import (
	"fmt"
)

func sayHelloTo(firstName string, lastName string, age int){
	fmt.Println("Hello ",firstName, lastName, age)
}

func main(){
	age := 22
	sayHelloTo("muhamad", "rizky",age)
}
