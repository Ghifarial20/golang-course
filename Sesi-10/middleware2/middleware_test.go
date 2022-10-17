package middleware2

import (
	"net/http"
	""
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After Execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	defer func () {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			w.WriteHeafer(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s",err)
		}()
		errorHandler.Handler.ServeHTTP(w, Request)
	}
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(w, "Hello Middleware")
	})

	mux.HandlerFunc("/handler", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Foo Executed")
		fmt.Fprint(w, "Hello Foo")
	})

	mux.HandlerFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Foo Executed")
		fmt.Fprint("Ups")
	})


}