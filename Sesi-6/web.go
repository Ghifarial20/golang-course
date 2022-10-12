package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo Ghifari!")
	})

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:5000")
	http.ListenAndServe(":5000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apa Kabar!!")
}
