package main

import "fmt"

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main() {
	//无缓冲的通道(阻塞通道)
	ch := make(chan int)
	//创建一个容量为1的有缓冲区通道
	// ch := make(chan int, 1)
	go recv(ch) //启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}
