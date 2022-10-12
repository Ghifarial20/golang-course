package main

import "fmt"

func main() {
	counter := 0
	name := "Ghifari"

	increment := func() {
		name := "Foo"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(counter)

	fmt.Println(name)
}
