package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//UserInfo 用户实体
type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析指定文件生成模板对象
	teml, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	//利用给定数据渲染模板 并将结果写入w
	user := UserInfo{
		Name:   "俊俊飞",
		Gender: "男",
		Age:    23,
	}
	//同理，当传入的变量是map时，也可以在模板文件中通过.根据key来取值。
	teml.Execute(w, user)
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed, err:", err)
		return
	}
}
