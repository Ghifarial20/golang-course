package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `
	{
		"full_name": "GhifariAL",
		"email": "Lustiansyah@gmail.com",
		"age": 22
	}`

	var temp interface{}

	err := json.Unmarshal([]byte(jsonString), &temp)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result := temp.(map[string]interface{}) // decode json ke interface

	fmt.Println("full_name:", result["full_name"])
	fmt.Println("email:", result["email"])
	fmt.Println("age:", result["age"])
}
