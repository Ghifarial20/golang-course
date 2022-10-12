package main

import (
	"fmt"
)

type Blacklist func(string) bool

func main() {
	// name := "Foo"
	// address := "Jakarta"
	// returnHello := getHello(name, address)
	firstName, _ := getFullName()
	months := []string{"Jan", "Feb", "Mar"}
	sliceMonths := getMonths(months)
	fmt.Println(sliceMonths)
	// for i := 0; i < 10; i++ {
	// 	// sayHi("Ghifari")
	// 	fmt.Println(returnHello)
	// }
	fmt.Println(getCompleteName())
	fmt.Println(firstName)

	total := sumAll(10, 20, 30, 40)
	fmt.Println(total)

	goodBye := getGoodBye
	fmt.Println(goodBye("Foo"))

	sayHelloWithFilter("Ghifari", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("testing", filter)

	blacklist := func(name string) bool {
		return name == "testing"
	}

	registerUser("Ghifari", blacklist)
	registerUser("testing", func(name string) bool {
		return name == "testing"
	})
}

func sayHello() {
	fmt.Println("Hello World")
}

func sayHi(name string) {
	fmt.Println("Hello", name)
}

func getHello(name, address string) string {
	result := "Hello nama saya " + name + " tinggal di " + address

	return result
}

func getFullName() (string, string) {
	return "Ghifari", "Lustiansyah"
}

func getMonths(months []string) []string {
	result := []string{}

	for i := 0; i < 2; i++ {
		result = append(result, months[i])
	}
	return result
}

func getCompleteName() (firstName, midName, lastName string) {
	firstName = "Ghifari"
	midName = "Ahmad"
	lastName = "Lustiansyah"

	return firstName, midName, lastName
}

func sumAll(numbers ...int) int { // variadik argument
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func listString(strings ...string) string {
	var result string

	for _, singleString := range strings {
		result += singleString
	}
	return result
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "testing" {
		return "..."
	}
	return name
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are Blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}
