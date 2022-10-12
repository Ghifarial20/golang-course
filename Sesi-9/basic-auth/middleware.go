// package main

// import "net/http"

// const USERNAME = "spiderman"
// const PASSWORD = "secret"

// func Auth(w http.ResponseWriter, r *http.Request) bool {
// 	username, password, ok := r.BasicAuth()
// 	if !ok {
// 		w.Write([]byte(`Something wrong`))
// 		return false
// 	}

// 	isValid := (username == USERNAME) && (password == PASSWORD)
// 	if !isValid {
// 		w.Write([]byte(`username/password is wrong`))
// 		return false
// 	}

// 	return true
// }

// func AllowOnlyGet(w http.ResponseWriter, r *http.Request) bool {
// 	if r.Method != "GET" {
// 		w.Write([]byte(`Only GET is Allowed`))
// 		return false
// 	}
// 	return true
// }
