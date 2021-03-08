package main

import "github.com/gin-gonic/gin"

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	// GET请求
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 启动HTTP服务　默认再127.0.0.1:9090启动服务
	r.Run()
}
