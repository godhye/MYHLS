package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf(req.URL.Path + "\n")
		p := "." + req.URL.Path
		if p == "./" {
			p = "./data/home.html"
		}
		http.ServeFile(w, req, p)
	})

	http.Handle("/data", http.FileServer(http.Dir(".")))

	http.HandleFunc("/hls", func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf(req.URL.Path)
		p := "." + req.URL.Path
		if p == "./hls" {
			p = "./data/test.html"
		}
		http.ServeFile(w, req, p)
	})
	http.ListenAndServe(":5000", nil)
}
