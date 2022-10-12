package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var eko Customer
	eko.Name = "Eko Sujono"
	eko.Address = "Jakarta"
	eko.Age = 17

	jono := Customer{
		Name:    "Jono Santuy",
		Address: "Malang",
		Age:     21,
	}

	fmt.Println(eko)
	fmt.Println(jono)

	eko.sayHi("Woy")
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello ", name)
}
