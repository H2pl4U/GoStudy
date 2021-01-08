package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	//使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type）
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", v.Name(), v.Kind())
}

func main() {
	var a float32 = 3.14
	reflectType(a) //type:float32
	var b int64 = 100
	reflectType(b) //type:int64
	var c *float64 //指针
	var d myInt    //自定义类型
	var e rune     //类型别名
	reflectType(c) //type: kind:ptr
	reflectType(d) //type:myInt kind:int64
	reflectType(e) //type:int32 kind:int32

	type person struct {
		name string
		age  int
	}

	type book struct{ title string }
	var f = person{
		name: "junhui",
		age:  22,
	}
	var g = book{title: "芜湖起飞"}
	reflectType(f) //type:person kind:struct
	reflectType(g) //type:book kind:struct
}
