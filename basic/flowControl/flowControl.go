package main

import "fmt"

func main() {
	//流程控制是每种编程语言控制逻辑走向和执行次序的重要部分，流程控制可以说是一门语言的“经脉”
	//if条件判断
	ifDemo1()
	ifDemo2()
	//for循环
	forDemo1()
	forDemo2()
	forDemo3()
	//for range(键值循环)
	//	1.数组、切片、字符串返回索引和值。
	//	2.map返回键和值。
	//	3.通道（channel）只返回通道内的值。

	//switch case
	switchDemo1()
	switchDemo2()
	switchDemo3()
	switchDemo4()

	//goto(跳转到指定标签)
	gotoDemo1()
	gotoDemo2()

	//break
	breakDemo()

	//continue
	continueDemo()

	//打印九九乘法表
	print99()
}

func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func forDemo1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func switchDemo1() {
	num := 3
	switch num {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("输入无效")
	}
}

func switchDemo2() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

func switchDemo3() {
	age := 23
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}

//fallthrough语法可以执行满足条件的case的下一个case
func switchDemo4() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

func gotoDemo1() {
	var flag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				//退出标签
				flag = true
				break
			}
			fmt.Printf("%v-%v \n", i, j)
		}
		//外层for循环判断
		if flag {
			break
		}
	}
}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakTag
			}
			fmt.Printf("%v-%v \n", i, j)
		}
	}
	return
breakTag:
	fmt.Println("结束for循环")
}

func breakDemo() {
BREAKDEMO:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO
			}
			fmt.Printf("%v-%v \n", i, j)
		}
	}
	fmt.Println("...")
}

func continueDemo() {
forloop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 && i == 2 {
				continue forloop
			}
			fmt.Printf("%v-%v \n", i, j)
		}
	}
}

func print99() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%v * %v = %-2v  ", j, i, i*j)
		}
		fmt.Println()
	}
}
