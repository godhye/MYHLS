package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("https://www.naver.com/")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//결과 출력
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(data))
}
