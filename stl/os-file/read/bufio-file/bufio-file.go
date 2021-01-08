package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file failed,err", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err", err)
			return
		}
		fmt.Print(line)
	}

	//ioutil.ReadFile读取整个文件
	content, err := ioutil.ReadFile("./bufio-file.go")
	if err != nil {
		fmt.Println("read file failed,err:", err)
		return
	}
	fmt.Println(string(content))
}
