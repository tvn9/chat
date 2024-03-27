package main

import "net/http"

func err(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	thread, err := data.Thread()
	if err != nil {
		error_message(w, r, "Connot get threads")
	} else {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, thread, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, thread, "layout", "private.navbar", "index")
		}
	}
}
