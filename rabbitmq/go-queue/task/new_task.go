package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func main() {
	//1.尝试连接RabbitMQ 建立连接
	//该连接抽象了套接字连接，并为我们处理协议版本协商和认证等
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("connect to RabbitMQ failed, err:%v\n", err)
		return
	}
	defer conn.Close()

	//2.创建一个通道，大多数API都是用过该通道操作
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("open a channel failed, err:%v\n", err)
		return
	}
	defer ch.Close()

	//3.声明队列
	q, err := ch.QueueDeclare(
		"task_queue", //name
		true,         //持久化
		false,        //delete when unused
		false,        //独有
		false,        //no-wait
		nil,          //arguments
	)
	if err != nil {
		fmt.Printf("declare a queue failed, err:%v\n", err)
		return
	}

	//4.将消息发布到声明到队列
	body := bodyFrom(os.Args)
	err = ch.Publish(
		"",     //exchange
		q.Name, //routing key
		false,  //立即
		false,  //强制
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, //持久
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		fmt.Printf("publish a message failed, err:%v\n", err)
		return
	}
	log.Printf(" [x] Sent %s", body)
}

// bodyFrom 从命令行获取将要发送的消息内容
func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello, jjh"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
