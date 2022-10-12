package main 

import(
	"fmt"
)

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("ini perhitungan ke-",counter)
	// 	counter++
	// }

	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	silceCountry := []string {"USA", "Canada", "Italy"}

	for j := 0; j < 3; j++ {
		fmt.Println(silceCountry[j])
	}
}