package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts") //menerima satu parameter dengan tipe data string yang merupakan sebuah urluntuk request tujuannya.
	if err != nil {
		log.Fatal(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(res.Body) // mengubah nilai yang kita akses dari field Body menjadi sebuah nilai dengan tipe data slice of byte
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}
