package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `
	{
		"full_name": "GhifariAL",
		"email": "Lustiansyah@gmail.com",
		"age": 22
	}`

	var result Employee

	err := json.Unmarshal([]byte(jsonString), &result) // decode data JSON ke Struct

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name:", result.FullName)
	fmt.Println("email:", result.Email)
	fmt.Println("age:", result.Age)
}
