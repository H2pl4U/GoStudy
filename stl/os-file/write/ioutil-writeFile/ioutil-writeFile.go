package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "hello world"
	err := ioutil.WriteFile("./test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
