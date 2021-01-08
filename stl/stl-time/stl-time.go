package main

import (
	"fmt"
	"time"
)

//时间类型
func timeDemo() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v \n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//时间戳
func timestampDemo() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

//使用time.Unix()函数可以将时间戳转为时间格式
func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//时间间隔
const (
	Nanosecond  = 1
	Microsecond = 1000 * Nanosecond
	Millisecond = 1000 * Microsecond
	Second      = 1000 * Millisecond
	Minute      = 60 * Second
	Hour        = 60 * Minute
)

func main() {
	timeDemo()
	timestampDemo()
	timestampDemo2(1606468169)

	//时间操作
	now := time.Now()
	later := now.Add(time.Hour) //当前时间加1个小时
	fmt.Println(later)
	//Add() Sub() Equal() Before() After()

	// tickDemo()
	formatDemo()

	timeStr := "2020/11/27 17:25:40"
	strToTime(timeStr)
}

// 定时器
func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

//时间格式化
func formatDemo() {
	now := time.Now()
	//格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	//24小时
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	//12小时
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02 "))
	fmt.Println(now.Format("2006/01/02"))
}

//解析字符串格式的时间
func strToTime(timeStr string) {
	now := time.Now()
	fmt.Println(now)
	//加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	//按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)          //输出时间
	fmt.Println(timeObj.Sub(now)) //与当前时间之差
}
