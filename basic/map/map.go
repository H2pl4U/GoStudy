package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//map的定义如下 map[KeyType]ValueType

	//map基本使用
	scoreMap := make(map[string]int, 6)
	scoreMap["热巴"] = 99
	scoreMap["娜扎"] = 98
	scoreMap["马飞"] = 59
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["马飞"])
	fmt.Printf("type of scoreMap:%T\n", scoreMap)

	userInfo := map[string]string{
		"username": "jjh",
		"password": "123",
	}
	fmt.Println(userInfo)

	//判断某个键是否存在 value, ok := map[key]
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok == true {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	//map的遍历 遍历map时的元素顺序与添加键值对的顺序无关
	//同时遍历键值对
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	//只遍历key
	for k := range scoreMap {
		fmt.Println(k)
	}

	//使用delete()函数删除键值对
	delete(scoreMap, "马飞")
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	//按照指定顺序遍历map
	rand.Seed(time.Now().UnixNano())
	var randMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("rand%02d", i)
		value := rand.Intn(100)
		randMap[key] = value
	}

	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range randMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, randMap[key])
	}

	//元素为map类型的切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "Kobe"
	mapSlice[0]["password"] = "0824"
	mapSlice[0]["address"] = "Laker"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	//值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)

}
