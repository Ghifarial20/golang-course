package main

import(
	"fmt"
	"strings"
)

func main() {
	var names [3]string // Buat array index

	names[0] = "Daniel"
	names[1] = "Ghifari"
	names[2] = "Lalo"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(strings.Repeat("#", 25))

	var values = [3]int {60,70,80}
	fmt.Println(values)
	fmt.Println(strings.Repeat("#", 25))
	
	countries := [3][2]string {
		{"Indonesia", "Singapore"},
		{"Australia", "USA"},
		{"Italy", "France"}}
	fmt.Println(len(countries))
	fmt.Println(strings.Repeat("#", 25))

	for i, listCountry := range countries {
		fmt.Println("Index ke ", i, listCountry)
		for _, country := range listCountry {
			fmt.Println(country)
		}
	}
	fmt.Println(strings.Repeat("#", 25))

	cities := [...]string {"Jakarta", "Solo", "Samarinda", "Pekanbaru"}

	fmt.Println(cities)
	fmt.Println(len(cities))

	fmt.Println(strings.Repeat("#", 25))

	// len(array)
	// array[index]
	// array[index] = values


}