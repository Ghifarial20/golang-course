package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Daniel",
		"email": "daniel@gmail.com",
	})

	respBody := bytes.NewBuffer(postBody) // untuk mengubah tipe data dari data yang ingin kita kirim menjadi interface io.Reader

	resp, err := http.Post("https://postman-echo.com/post", "application/json", respBody)

	if err != nil {
		log.Fatalf("An error occured %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
