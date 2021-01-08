package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//GET请求的参数需要使用Go语言内置的net/url这个标准库来处理。

func main() {
	apiURL := "https://open.ys7.com/api/lapp/device/list"
	//参数
	data := url.Values{}
	data.Set("accessToken", "at.5b7bfafi0qlwfj6dc942f2r4dm2yjwq7-838h19oboq-1d9jen3-ac8yblgzg")
	u, err := url.ParseRequestURI(apiURL)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode() //URL encode
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
