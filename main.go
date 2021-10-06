package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	_ "strings"
)

type Person struct {
		Name  string
		Phone string
	}

func main() {

	/*  reqbody := bytes.NewBufferString("Post Plain text")  1
	resp, err := http.Post("http://httpbin.org/post", "text/plain", &reqbody)
	if err != nil {
		panic(err)
	}
	*/

	//io.reader : 어떤 구조체든 매개변수로 바이트 슬라이스를 받고,
	//정수와 에러 값을 리턴하는 Read 함수를 가지고 있으면 io.Reader 인터페이스를 따른다고 할 수 있습니다

	//var buf bytes.Buffer //힙메모리 할당 발생
	//buf.WriteString("name")
	//resp, err := http.Post("http://httpbin.org/post", "text/plain", &buf)  2

	//buf := strings.NewReader("age")  3

	// http.post 방식
	buf := bytes.NewReader([]byte{13, 1, 2, 3}) //4
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

	//http.PostForm
	res, err := http.PostForm("http://httpbin.org/post", url.Values{"name": {"eee"}, "age": {"33"}})
	if err == nil {
		str := string(respBody)
		println(str)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err == nil {
		str := string(body)
		println(str)
	}

	
	//json 전송
	person := Person{"name", "0109-"}
	pbytes, err := json.Marshal(person)
	if err == nil {
		str := string(pbytes)
		println("json", str)

	}
	buff := bytes.NewBuffer(pbytes)

	r, err := http.Post("http://httpbin.org/post", "appication/json", buff)

	body1, err := ioutil.ReadAll(r.Body)
	if err == nil {
		str := string(body1)
		println(str)
	}

	//person1 := Person{"dicik", "3300-3"}
	//pbytes1, _ := xml.Marshal(person1)
	pbytes1, _ := xml.Marshal([]string{"a", "b", "c"})
	buff2 := bytes.NewBuffer(pbytes1)

	r2, err := http.Post("http://httpbin.org/post", "appication/xml", buff2)

	body2, err := ioutil.ReadAll(r2.Body)
	if err == nil {
		str := string(body2)
		println(str)
	}

}
