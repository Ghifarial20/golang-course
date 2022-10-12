package main

import (
	"errors"
	"fmt"
)

func main() {
	// var name string
	// name = "Ghifari AL"

	// anotherName := "G A L"

	// varName, age, country := "GAL", 30, "Indonesia"

	// fmt.Println(name)
	// fmt.Println(anotherName)
	// fmt.Println(varName, age, country)
	// forInt()
	// dataType()

	anotherAge := 35
	tes, err := returnError(anotherAge)
	fmt.Println(tes, err)
}

func forInt() {
	sliceCountry := []string{"Indonesia", "Malaysia", "Singapore"}

	for _, singleCountry := range sliceCountry {
		fmt.Println(singleCountry)
	}
}

func dataType() {
	var age int8    // tipe data interger
	var result bool // tipe data boolean defaultnya FALSE

	newCountry := "Indonesia"
	age = 17

	if result {
		newCountry = "Malaysia"
	}

	fmt.Println(age)
	fmt.Println(result)
	fmt.Println("Hasil if adalah ", newCountry)
}

func returnError(anotherAge int) (bool, error) {
	if anotherAge > 30 {
		return false, errors.New("age more than 30")
	}
	return true, nil
}
