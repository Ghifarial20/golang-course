package main

import (
	"fmt"
)

const TEST_CONST int = 1

func main() {
	fmt.Println(TEST_CONST)

	fmt.Println(returnError())
}

func returnError() int {
	return TEST_CONST
}
