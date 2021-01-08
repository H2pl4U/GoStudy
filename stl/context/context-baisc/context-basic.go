package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exit bool

//基本示例
// func worker1() {
// 	for {
// 		fmt.Println("worker")
// 		time.Sleep(time.Second)
// 	}
// 	// 如何接收外部命令实现退出
// 	wg.Done()
// }

//全局变量方式
func worker2() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	wg.Done()
}

//通道方式
func worker3(exitChan chan struct{}) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-exitChan: //等待接收上级通知
			break LOOP
		default:
		}
	}
}

//官方版方案
func worker4(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

//当子goroutine又开启外一个goroutine时，只需要将ctx传入即可
func woker5(ctx context.Context) {
	go worker6(ctx)
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): //等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func worker6(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
}

func main() {
	// wg.Add(1)
	// go worker1()
	// //优雅的结束goroutine
	// wg.Wait()
	// fmt.Println("over")

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker4(ctx)
	time.Sleep(time.Second * 3)
	cancel() //通知子goroutine结束
	wg.Wait()
	fmt.Println("over")

}
