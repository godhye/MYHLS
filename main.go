package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.HandleFunc("/bye", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("bye world  with ," + req.URL.RawQuery))
	})

	http.ListenAndServe(":5000", nil)
}
