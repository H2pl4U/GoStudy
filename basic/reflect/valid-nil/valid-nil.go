package main

import (
	"fmt"
	"reflect"
)

//IsNil()报告v持有的值是否为nil。
//v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。

//IsValid()返回v是否持有一个值。
//如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。

func main() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	//实例化一个匿名结构体
	b := struct{}{}
	//尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员", reflect.ValueOf(b).FieldByName("abc").IsValid())
	//尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	//尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键:", reflect.ValueOf(c).MapIndex(reflect.ValueOf("热巴")))
}
