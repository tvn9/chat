package main

import "net/http"

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := data.User

}
