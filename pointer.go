package main

import "fmt"

// Address is ...
type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address){ // pointer di function
	address.Country = "Jerman"
}

func main() {
	address1 := Address{"Tegal", "Jawa Tengah", "Indonesia"}
	// var address4 *Address = &Address{"Tegal", "Jawa Tengah", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Banyumas"

	// address2 = &Address{"Semarang", "Jawa Tengah", "Indonesia"} // membuat lokasi memori baru
	*address2 = Address{"Surabaya", "Jawa Timur", "Indonesia"} // jika menggunakan tanda bintang didepanya maka siapapun yg mengacu pada memori (struct Address) datanya akan ikut berubah termasuk address 1,2,3

	fmt.Println(address1) // address1 tidak berubah
	fmt.Println(address2)
	fmt.Println(address3)


	var address4 *Address = new(Address)
	address4.City = "Bandung"
	fmt.Println(address4)

	var alamat = Address{
		City: "Purwokerto",
		Province: "Jawa Tengah",
		Country: "",
	}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
