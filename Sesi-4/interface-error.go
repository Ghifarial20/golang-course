package main

import (
	"errors"
	"fmt"
)

func main() {
	hasil, err := Pembagi(100, 200)

	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}

func Pembagi(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Pembangi tidak boleh 0")
	} else {
		result := x / y
		return result, nil
	}
}
