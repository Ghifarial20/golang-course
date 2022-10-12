package main 

import(
	"fmt"
)

func main() {
	countries := []string {"Indonesia", "Brunei", "Malaysia"}

	fmt.Println(countries)

	months := [12]string {"Januari", "Februari", "Maret", "April", "Mei",
	"Juni", "Juli", "Agustus", "Sept", "Okt", "Nov", "Des"}

	slice := months[4:7]

	fmt.Println(slice)
	fmt.Println(len(slice))
	slice = append(slice, "Jakarta")
	fmt.Println(slice)

	newCity := []string {"Jakarta", "Solo", "Denpasar"}

	newSlice := make([]string, 2)
	newSlice[0] = "Ghifari"
	newSlice[1] = "Jhon"

	sliceBaru := make([]string, len(newSlice))

}