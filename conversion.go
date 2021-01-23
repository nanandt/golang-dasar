package main

import (
	"fmt"
)

func main(){
	var nilai32 int32 = 3276
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) // jika kapasitas tdk mencukupi maka niainya akan kembali ke titik awal

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Rizky"
	var e byte = name[3]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}
