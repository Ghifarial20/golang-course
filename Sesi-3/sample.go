package main

import (
	"fmt"
)

func main() {
	months := [12]string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	fmt.Println(cap(months))

	sliceMonths := months[3:12]

	fmt.Println(sliceMonths)
}
