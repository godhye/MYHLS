package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	_ "strings"
)

func main() {

	/*  reqbody := bytes.NewBufferString("Post Plain text")
	resp, err := http.Post("http://httpbin.org/post", "text/plain", &reqbody)
	if err != nil {
		panic(err)
	}
	*/

	//io.reader : 어떤 구조체든 매개변수로 바이트 슬라이스를 받고,
	//정수와 에러 값을 리턴하는 Read 함수를 가지고 있으면 io.Reader 인터페이스를 따른다고 할 수 있습니다

	//var buf bytes.Buffer //힙메모리 할당 발생
	//buf.WriteString("name")
	//resp, err := http.Post("http://httpbin.org/post", "text/plain", &buf)

	//buf := strings.NewReader("age")
	buf := bytes.NewReader([]byte{13, 1, 2, 3})
	resp, err := http.Post("http://httpbin.org/post", "text/plain", buf)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}

}
