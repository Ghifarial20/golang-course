package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var student = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudent() []*Student { return student }

func SelectStudent(id string) *Student {
	for _, each := range student {
		if each.Id == id {
			return each
		}
	}
	return nil
}

func init() {
	student = append(student, &Student{Id: "s001", Name: "Ghifari", Grade: 2})
	student = append(student, &Student{Id: "s002", Name: "Ahmad", Grade: 3})
	student = append(student, &Student{Id: "s003", Name: "Lustiansyah", Grade: 4})
}

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":8080"

	fmt.Println("server started at localhost:8080")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGet(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudent())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

const USERNAME = "spiderman"
const PASSWORD = "secret"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`Something wrong`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte(`username/password is wrong`))
		return false
	}

	return true
}

func AllowOnlyGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte(`Only GET is Allowed`))
		return false
	}
	return true
}
