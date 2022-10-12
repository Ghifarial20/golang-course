package main

import(
	"fmt"
)

func main(){
	for i := 0; i < 10; i++ {
		if i == 2 {
			fmt.Println(i)
			break
		}
		fmt.Println("Perulangan ke-", i)
	}
}