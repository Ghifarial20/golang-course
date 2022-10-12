package Project

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println()
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("cek")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
