package main

import (
	"fmt"

	m "h2pl4u.com/basic/package"
)

//全局变量定义标准 var 变量名 类型
var name string = "name" //""
var age int = 22         //0
var sex bool = true      //false

//常量
const constNum = 111

//批量常量
const (
	pi = 3.1415
	e  = 2.7182
)

//iota go语言的常量计数器，只能在常量的表达式中使用
//在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次
const (
	p1 = iota //0
	p2        //1
	p3        //2
	p4        //3
)

// const (
// 	p1 = iota	//0
// 	p2 = 100	//100
// 	p3			//2
// 	p4			//3
// )

//使用_跳过一些值
// const (
// 	p1 = iota	//0
// 	p2			//1
// 	_
// 	p4			//3
// )

//批量定义变量
// var (
// 	name string
// 	age  int
// 	sex  bool
// )

func main() {
	//局部变量 标准语法
	var str1 string = "str1"
	var num1 int = 1
	var isBool1 bool = true

	//类型推导
	var str2 = "str2"
	var num2 = 2
	var isBool2 = false

	//简短变量声明 在函数内部，可以使用更简略的 := 方式声明并初始化变量
	str3 := "str3"
	num3 := 3
	isBool3 := true

	//匿名变量 在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。 匿名变量用一个下划线_表示
	fmt.Println("----------匿名变量-------------")
	x, _ := foo()
	_, y := foo()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("----------常量-----------------")
	fmt.Println(constNum)
	fmt.Println(pi)
	fmt.Println(e)
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)
	fmt.Println("----------全局变量-------------")
	fmt.Printf("name:%s,", name)
	fmt.Printf("age:%d,", age)
	fmt.Println(sex)
	fmt.Println("----------局部变量-------------")
	fmt.Printf("str1:%s,", str1)
	fmt.Printf("num1:%d,", num1)
	fmt.Println(isBool1)
	fmt.Println("----------类型推导-------------")
	fmt.Printf("str2:%s,", str2)
	fmt.Printf("num2:%d,", num2)
	fmt.Println(isBool2)
	fmt.Println("--------简短变量声明-----------")
	fmt.Printf("str3:%s,", str3)
	fmt.Printf("num3:%d,", num3)
	fmt.Println(isBool3)

	//调用package
	sum := m.Add(4654, m.Mode)
	fmt.Println(sum)
}

func foo() (int, string) {
	return 23, "bar"
}
