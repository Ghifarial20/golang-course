package simple_api

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!!")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestServerMux(t *testing.T) {
	mux := http.NewServerMux()

	mux.HandlerFunc("/", func(w http.ResponseWriter, r http.Request) {
		fmt.Fprintf(w, "Hello World!!")
	})

	mux.HandlerFunc("/hi", func(w http.ResponseWriter, r http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	mux.HandlerFunc("/images/", func(w http.ResponseWriter, r http.Request) {
		fmt.Fprintf(w, "Image")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
