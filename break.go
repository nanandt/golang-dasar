package main

import "fmt"

func main(){
	for i := 0; i < 10; i++{
		if i == 5{
			break // menghentikan paksa perulangan dengan paksa jika suatu kondisi terpenuhi dan kode dibawahnya tdk dijalankan
		}
		fmt.Println("perulangan ke ", i)
	}
}
