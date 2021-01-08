package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Atoi 将string整数转换成int
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}

	//Itoa 将int类型转换成对应的字符串
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"

	//parse系列函数
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	i, err := strconv.ParseInt("-2", 10, 64)
	u, err := strconv.ParseUint("2", 10, 64)
	fmt.Println(b, f, i, u)

	//Format系列函数
	str1 := strconv.FormatBool(true)
	str2 := strconv.FormatFloat(3.165423, 'E', -1, 64)
	str3 := strconv.FormatInt(-2, 16)
	str4 := strconv.FormatUint(2, 16)
	fmt.Println(str1, str2, str3, str4)
}
