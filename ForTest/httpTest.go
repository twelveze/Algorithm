package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://golang.org", nil)
	if err != nil{
		fmt.Println(err)
	}

	req.Close = true
	//req.Header.Add("Connection", "close") // 等效的关闭方式

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}