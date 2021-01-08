package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//只读方式打开当前目录下的main.go文件
	file, err := os.Open("./os-file.go")
	//关闭文件
	defer file.Close()
	if err != nil {
		fmt.Println("open file failed! err:", err)
		return
	}
	//使用Read方法读取数据 file.Read() byte数组为读取文件的大小 单位是字节
	var tmp = make([]byte, 256)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))
	//循环读取文件
	var content []byte
	var tmp2 = make([]byte, 128)
	for {
		n, err := file.Read(tmp2)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp2[:n]...)
	}
	fmt.Println(string(content))
}
