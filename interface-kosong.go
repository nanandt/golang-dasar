package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return 2
	}
	return "ups salah"
}

func main() {
	kosong := Ups(4)
	fmt.Println(kosong)
}
