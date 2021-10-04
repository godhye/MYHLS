package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//get으로 resp에 데이터 가져옴

	fmt.Println("Get Method : http Get 사용 / naver.com")
	resp, err := http.Get("https://www.naver.com/")
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n\n", string(data)[0:100])

	//data 초기화
	data = []byte{}

	fmt.Println("Get Method : Request 객체, Http.Clinet 사용 /  google.com")
	//request 객체 생성
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	//헤더 추가
	req.Header.Add("User-Agent", "Crawler")

	client := &http.Client{}

	//Client 객체를 통해 req 호출
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	//결과 출력
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n\n", string(data)[0:100])
}
