package main

import "fmt"

func main() {

	name := "rizky"
	counter := 0

	increment := func() {
		name := "wahyu"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
