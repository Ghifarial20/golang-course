package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

func main() {
	sample := Sample{"Ghifari"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)

	sample.Name = ""
	fmt.Println(IsValid(sample))

}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}
