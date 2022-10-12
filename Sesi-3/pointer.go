package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Semarang", "Jawa Tengah", "Indondesia"}

	address2 := address1
	address3 := &address1

	address2.City = "Bandung"
	address3.City = "Solo"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = &address1
	var address5 *Address = &address1
	fmt.Println(address4)
	fmt.Println(address5)

	var alamat1 *Address = new(Address)
	alamat1.City = "Jakarta"
	fmt.Println(alamat1)

	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}

	// fmt.Println(alamat)

	var alamatPointer *Address = &alamat
	ChangeCountryIndonesia(alamatPointer)
	fmt.Println(alamatPointer)
}

func ChangeCountryIndonesia(address *Address) {
	address.Country = "Indonesia"
}
