package main

import(
	"fmt"
)

func main() {
	result := false

	if (5 % 2 != 0) {
		result = true
		fmt.Println(result)
	}
	fmt.Println(result)
}