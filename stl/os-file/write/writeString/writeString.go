package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	str := "强啊鸡飞\n"
	file.Write([]byte(str))  //写入字节切片数据
	file.WriteString("弱啊巴蒂") //直接写入字符串数据
}
