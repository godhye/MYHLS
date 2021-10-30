package main

import (
	"fmt"
	"net/http"
)

func main() {
	//127.0.0.1:5000/
	//메인 홈페이지 , homehtml 에는 사진 , 사진위에 글씨 출력
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf(req.URL.Path + "\n")
		p := "." + req.URL.Path
		if p == "./" {
			p = "./data/home.html"
		}
		//p에 있는 파일을 w에 담아서 보냄
		http.ServeFile(w, req, p)
	})

	//127.0.0.1:5000/data ./data폴더에 있는 파일 다운로드 서버 제공
	http.Handle("/data", http.FileServer(http.Dir(".")))

	//127.0.0.1:5000/hls  버튼 눌러 동영상 재생 , hls 제공 예정
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
