package main

import "fmt"

func main() {
	sample := "mala"
	fmt.Println(Palindrome(sample))
}

func Palindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		if string(str[i]) != string(str[len(str)-1-i]) {
			return false
		}
	}
	return true
}
