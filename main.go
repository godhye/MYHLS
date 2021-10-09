package main

import (
	"net/http"
)

type testHandler struct {
	http.Handler
}

func (h *testHandler) ServeHTTP(W http.ResponseWriter, req *http.Request) {
	str := "Request path is " + req.URL.Path + "\n"
	str += "Parameter    is " + req.URL.RawQuery + "\n"
	W.Write([]byte(str))
}
func main() {

	http.Handle("/bye", new(testHandler))

	http.ListenAndServe(":5000", nil)
}
