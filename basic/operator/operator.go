package main

import (
	"fmt"
	"strconv"
)

var (
	a int = 1
	b int = 2
	c int = 3
	d int = 4
)

func main() {
	//运算运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(float32(a) / float32(b))
	fmt.Println(d % c)
	//关系运算符
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
	//逻辑运算符
	fmt.Println(a < c && b > a)
	fmt.Println(a == c || b > c)
	fmt.Println(!(a > b))
	//位运算符
	fmt.Println(c & b)
	fmt.Println(c | b)
	fmt.Println(c ^ b)
	fmt.Println(c >> b)
	fmt.Println(c << b)
	//赋值运算符
	var num int
	num = a
	num += a
	num -= a
	num *= d
	num /= b
	num %= c
	num >>= 1
	num <<= 1
	num &= 1
	num |= 1
	num ^= 2
	fmt.Println("---------------")
	index := index("1234321")
	fmt.Println(index)
}

//寻找只出现过一次的数字
func index(str string) int {
	var arr [64]int
	for i := 0; i < len(str); i++ {
		num, err := strconv.Atoi(string(str[i]))
		if err == nil {
			arr[num]++
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			return i
		}
	}
	return -1
}
