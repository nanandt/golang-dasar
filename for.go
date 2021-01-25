package main

import "fmt"

func main(){

	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke ",counter)
		counter++
	}

	fmt.Println("")


	for number := 1; number <= 10; number++{
		fmt.Println("angka ke", number)
	}

	slice := []string{"muhamad", "fatih", "rizky"}

	biodata := make(map[string]string)
		biodata["name"] = "rizky"
		biodata["address"] = "tegal"


	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}

	for _, value := range slice{ // _ akan memberitahu golang bahwa variable tsb tdk dibutuhkan
		// fmt.Println("index ke ",i, "=", value)
		fmt.Println(value)
	}

	for key, value := range biodata{
		fmt.Println(key, "=", value)
	}

}


// - number := 1 adalah init statement dan akan dieksukesi sekali saja sebelum perulangan dieksekusi
// - number++ adalah post statement dan akan terus dieksekusi sampai kondisinya terpenuhi
//   walaupun letakanya tidak didalam {}
