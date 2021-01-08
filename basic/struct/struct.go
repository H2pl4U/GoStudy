package main

import (
	"encoding/json"
	"fmt"
)

//自定义类型  Go语言中可以使用type关键字来定义自定义类型,如string、整型、浮点型、布尔等数据类型
//通过type关键字的定义，MyInt就是一种新的类型，它具有int的特性
type myType int

//类型别名 类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型  如type byte = int8
type intAlias = int

// 结构体的定义
//结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）
type student struct {
	id, classID int
	name        string
	sex         bool
}

// 方法和接收者
// Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。
// 接收者的概念就类似于其他语言中的this或者 self student构造函数
func newStudent(id int, classID int, name string, sex bool) *student {
	return &student{
		id:      id,
		classID: classID,
		name:    name,
		sex:     sex,
	}
}

// 值类型的接收者
func (stu student) Dream() {
	fmt.Printf("%s的梦想是当一名架构师", stu.name)
	fmt.Println()
}

//任意类型添加方法
func (m myType) SayHello() {
	fmt.Println("Hello,定义了myType(int)")
}

// 结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
type part struct {
	string
	int
}

type teacher struct {
	ID   int `json:"id"` //通过指定tag实现json序列化该字段时的key
	Name string
}

//嵌套结构体
type class struct {
	ClassName string
	Student   student
	Teacher   teacher
}

// 结构体的“继承”
type animal struct {
	name string
}

func (a *animal) speed() {
	fmt.Printf("%s的速度40km/h\n", a.name)
}

type dog struct {
	age     int
	*animal //通过嵌套匿名结构体实现继承
}

func (d *dog) wang() {
	fmt.Printf("%s会汪汪汪\n", d.name)
}

func main() {
	var num1 myType = 1
	fmt.Printf("%T\n", num1) //main.myType
	fmt.Printf("%d\n", num1)
	var num2 intAlias = 2
	fmt.Printf("%T\n", num2) //int
	fmt.Printf("%d\n", num2)

	//实例化结构体
	var stu1 student
	stu1.id = 1
	stu1.classID = 1001
	stu1.name = "jjh"
	stu1.sex = true
	fmt.Println(stu1)

	//匿名结构体
	var class01 struct {
		classID   int
		className string
	}
	class01.classID = 2001
	class01.className = "二年级一班"
	fmt.Println(class01)

	//创建指针类型结构体
	//通过使用new关键字对结构体进行实例化，得到的是结构体的地址
	var stu2 = new(student)
	fmt.Printf("stu2=%#v\n", stu2) //stu2=&main.student{id:0, classID:0, name:"", sex:false}
	stu2.id = 2
	stu2.classID = 2001
	stu2.name = "xiaoli"
	stu2.sex = true
	fmt.Printf("stu2=%#v\n", stu2) //stu2=&main.student{id:2, classID:2001, name:"xiaoli", sex:true}

	//取结构体的地址实例化
	//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
	stu3 := &student{}
	fmt.Printf("stu3=%#v\n", stu3) //stu3=&main.student{id:0, classID:0, name:"", sex:false}
	stu3.id = 3
	stu3.classID = 2002
	stu3.name = "xiaoliu"
	stu3.sex = false
	fmt.Printf("stu3=%#v\n", stu3) //stu3=&main.student{id:3, classID:2002, name:"xiaoliu", sex:false}

	//使用键值对初始化
	stu4 := student{
		id:      4,
		classID: 1002,
		name:    "xiaozhang",
		sex:     false,
	}
	fmt.Printf("stu4=%#v\n", stu4) //stu4=main.student{id:4, classID:1002, name:"xiaozhang", sex:false}

	//也可以对结构体指针进行键值对初始化，例如：
	stu5 := &student{
		id: 5,
		//无classID可不写
		name: "xiaoma",
		sex:  true,
	}
	fmt.Printf("stu5=%#v\n", stu5) //stu5=&main.student{id:5, classID:0, name:"xiaoma", sex:true}

	//使用值的列表初始化,注意必须初始化结构体的所有字段且顺序一致，不能和键值初始化方式混用
	stu6 := student{
		6,
		2003,
		"xiaosong",
		true,
	}
	fmt.Printf("stu6=%#v\n", stu6) //stu6=main.student{id:6, classID:2003, name:"xiaosong", sex:true}
	//结构体占用一块连续的内存
	fmt.Printf("stu6.id %p\n", &stu6.id)
	fmt.Printf("stu6.classID %p\n", &stu6.classID)
	fmt.Printf("stu6.name %p\n", &stu6.name)
	fmt.Printf("stu6.sex %p\n", &stu6.sex)

	//方法和接收者 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型
	stu7 := newStudent(7, 3001, "小刘", false)
	stu7.Dream()

	var m1 myType
	m1.SayHello()
	m1 = 199
	fmt.Printf("%#v  %T\n", m1, m1) //199  main.myType

	//结构体的匿名字段
	p1 := part{
		"junfei",
		22,
	}
	fmt.Printf("%#v \n", p1)
	fmt.Println(p1.string, p1.int)

	//嵌套结构体
	csClass := class{
		ClassName: "计科一班",
		Student: student{
			id:      1,
			classID: 2,
			name:    "mff",
			sex:     true,
		},
		Teacher: teacher{
			ID:   2,
			Name: "Teacher马",
		},
	}
	csClass.Student.name = "lbw"
	csClass.Teacher.Name = "teacherma"
	fmt.Printf("csClass:%#v\n", csClass)

	//结构体的继承
	d1 := &dog{
		age: 2,
		animal: &animal{ //嵌套的是结构体指针
			name: "旺财",
		},
	}
	d1.wang()
	d1.speed()

	// json序列化时结构体熟悉首字母必须大写
	data1, _ := json.Marshal(csClass)
	fmt.Println(string(data1))
}
