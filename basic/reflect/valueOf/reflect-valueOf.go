package main

import (
	"fmt"
	"reflect"
)

//reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。
//reflect.Value与原始值之间可以互相转换。

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		//v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		//v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f \n", float32(v.Float()))
	case reflect.Float64:
		//v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float64, value is %f \n", float64(v.Float()))
	}
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本 reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	//反射中使用Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200) //修改的是副本 reflect包会引发panic
	}
}

func main() {
	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a) //type is float32,value is 3.140000
	reflectValue(b) //type is int64,value is 100
	//将int类型的原始值转换成reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c :%T \n", c) //type c:reflect.Value

	var d int64 = 100
	reflectSetValue2(&d)
	fmt.Println(d)
}
