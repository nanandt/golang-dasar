package main

import "fmt"

// Blacklist is ...
type Blacklist func(string) bool


func registerUser(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("You are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}


func main(){

	blacklist := func(name string) bool {
		return name == "hasan"
	}

	registerUser("Panji",blacklist)
	registerUser("hasan",blacklist)

	registerUser("joniii", func(name string) bool{
		return name == "joni"
	})

	registerUser("kenthu", func(name string) bool{
		return name == "joni"
	})

}
