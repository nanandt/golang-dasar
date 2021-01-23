package main

import (
	"fmt"
)

func main(){
	type noKtp string
	type Married bool


	var noKtpRizky noKtp = "2219718262"
	var marriedStatus Married = true

	fmt.Println(noKtpRizky)
	fmt.Println(marriedStatus)
}
