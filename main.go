package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type staticHandler struct {
	http.Handler
}

func getContentType(localpath string) string {
	contenttype := ""
	ext := filepath.Ext(localpath)

	switch ext {

	case ".html":
		contenttype = "text.html"
	case ".css":
		contenttype = "text.css"
	case ".js":
		contenttype = "text.js"
	case ".png":
		contenttype = "text.png"
	}

	return contenttype
}
func (h *staticHandler) ServeHTTP(W http.ResponseWriter, req *http.Request) {
	//파일 위치 읽기
	localPath := "data" + req.URL.Path
	content, err := ioutil.ReadFile(localPath)
	if err != nil {
		W.WriteHeader(404)
		W.Write([]byte(http.StatusText(404)))
		return
	}

	contentType := getContentType(localPath)

	//파일 형식 지정
	W.Header().Add("Content-Type", contentType)
	W.Write(content)
	//응답
}

type testHandler struct {
	http.Handler
}

//구조체 testHandler의 메소드 구현
func (h *testHandler) ServeHTTP(W http.ResponseWriter, req *http.Request) {
	//Write([]byte)를 이용해 바로 쓰는거 보다 깔끔하게 string에 담아서 쓰기
	str := "Request path is " + req.URL.Path + "\n"
	str += "Parameter    is " + req.URL.RawQuery + "\n"
	W.Write([]byte(str))
}

func main() {

	http.Handle("/bye", new(testHandler))
	http.Handle("/", new(staticHandler))

	http.ListenAndServe(":5000", nil)
}
