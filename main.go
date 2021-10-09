package main

import (
	"net/http"
	"net/url"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.HandleFunc("/bye", func(w http.ResponseWriter, req *http.Request) {

		//파라미터 출력을 위한 파라미터 파싱
		parameter, err := url.ParseQuery(req.URL.RawQuery)
		if err == nil {
			for k, v := range parameter {
				w.Write([]byte(k + " "))
				for i := 0; i < len(v); i++ {
					w.Write([]byte(v[i]))
				}

				w.Write([]byte("\n"))

			}
		}

		w.Write([]byte("bye world  with "))
	})

	http.ListenAndServe(":5000", nil)
}
