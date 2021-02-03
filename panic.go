package main

import "fmt"

func endApp() {

	message := recover() // recover disimpan di func defernya
	if message != nil {

	fmt.Println("Error dengan message", message)
	}

	fmt.Println("Aplikasi selesai")
}

func runApp(error bool){
	defer endApp()

	if error {
		panic("APLIKASI ERROR") // jika menggunakan panic makan defernya tetep jalan dan
								// dan functionya berhenti
	}
	fmt.Println("Aplikasi Berjalan")
}

func main(){
	runApp(true)
	fmt.Println("Aplikasi tetep jalan walopun ada error")
}
