package main

import(
	"fmt"
)

func main() {
	silceCountry := []string{"Indonesia", "Singapore", "Australia", "Myanmar"}

	var length int
	length = len(silceCountry)

	if length > 10 {
		fmt.Println("Jumlah negara terlalu banyak!!")
	} else if length < 10 && length > 3 {
		if silceCountry[0] == "Indonesia" {
			fmt.Println("Ini adalah negara kita tercinta")
		}
	} else {
		fmt.Println("Capaian tidak terpenuhi")}
}