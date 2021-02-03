package main

import "fmt"

/**
struct adalah representasi data di aplikasi kita / field nya
*/

// Customer is ..
type Customer struct {
	Name, Address string
	Age           int
	Married bool
}

func (customer Customer) sayHi(name string){
	fmt.Println("Hello,", name, "my name is", customer.Name)

}

func (a Customer) sayHuuu(){
	fmt.Println("Huuu from", a.Name)
}

func main() {

	var rizky Customer
	rizky.Name = "Muhamad fatikhurrizky"
	rizky.Address = "Indonesia"
	rizky.Age = 22
	rizky.sayHi("ahmad")

	rizky.sayHuuu()

	// fmt.Println(rizky)
	// fmt.Println(rizky.Name)
	// fmt.Println(rizky.Address)
	// fmt.Println(rizky.Age)

	// harun := Customer{
	// 	Name: "Harun",
	// 	Address: "Bekasi",
	// 	Age: 11,
	// 	Married: true,
	// }

	// fmt.Println(harun)

	// bayu := Customer{"Budi", "Tegal", 34, true}
	// fmt.Println(bayu)

}
