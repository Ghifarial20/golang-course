package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Ghifari Ahmad L"
		fmt.Print("Selesai\n")
	}()

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string) //pass by reference
	defer close(channel)

	abc := "string Kosong"

	go GiveMeResponse(channel, abc) // pass by value

	data := <-channel

	fmt.Println(data, abc)
	time.Sleep(2 * time.Second)
}

func GiveMeResponse(channel chan string, abc string) {
	time.Sleep(2 * time.Second)
	channel <- "HAHAHAHA"
	abc = "bukan string kosong"
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "GHIFARI"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
}

func TestDeclareChannel(t *testing.T) {
	channel := make(chan string)
	go Parameter(channel)
	data := <-channel
	fmt.Println(data)
}

func Parameter(channel chan string) {
	channel <- "Coba-coba"
}
