package main

import(
	"fmt"
)

func main() {
	person := map[string]string {
		"name": "Supri",
		"address": "Semarang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	food := make(map[string]string)

	food["pokok"] = "Nasi"
	food["snack"] = "Kue"
	food["dessert"] = "agar"

	fmt.Println(food)
	delete(food, "snack")
	fmt.Println(food)
}