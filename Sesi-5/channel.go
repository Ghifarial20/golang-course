package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Ghifari Ahmad L"
		fmt.Print("Selesai")
	}()

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
