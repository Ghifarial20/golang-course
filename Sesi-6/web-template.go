package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "Ghifari AL",
			"Message": "Have a Nice Day",
		}

		var t, err = template.ParseFiles("template.html")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
