package main

import (
	"fmt"
)

func main(){
	name := "wahyuuu"

	if name == "rizky"{
		fmt.Println("Hallo", name)
	} else if name == "wahyu"{
		fmt.Println("Hallo Wahyu")
	}else if name == "sugiono"{
		fmt.Println("Hallo Sugiono")
	}else {
		fmt.Println("Hai leh nalan?")
	}

	// var length = len(name)
	// if length > 5 {
	// 	fmt.Println("nama terlalu panjang")
	// } else {
	// 	fmt.Println("nama sudah benar")
	// }

	if length := len(name); length > 5{ // short statement
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
