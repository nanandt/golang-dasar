package main

import (
	"fmt"
)

func main(){
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "diubah"
	// fmt.Println(slice1)

	// slice1[0] = "mangkok"
	// fmt.Println(slice1)

	var slice2 = months[10:]
	// var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "garpu")
	fmt.Println(slice3)

	slice3[1] = "desember baru"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "rizky"
	newSlice[1] = "khuuuuur"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))


	copySlice := make([]string, len(newSlice), cap(newSlice))

	copy(copySlice,newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1,2}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)



}
