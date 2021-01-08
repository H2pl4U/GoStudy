package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	//浮点数
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	//复数
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	//符号转义
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")

	//换行
	s1 := `第一行
		   第二行
		   第三行`
	fmt.Println(s1)

	s2 := "a,b,c,d,e,f"
	//长度
	fmt.Println(len(s2))
	//切割
	fmt.Println(strings.Split(s2, ","))
	//包含
	fmt.Println(strings.Contains(s2, "b,c"))
	//前缀判断
	fmt.Println(strings.HasPrefix(s2, "a,b"))
	//后缀判断
	fmt.Println(strings.HasSuffix(s2, "e,f"))
	//子串出现的位置
	fmt.Println(strings.Index(s2, "c"))
	fmt.Println(strings.LastIndex(s2, ","))

	// 遍历字符串
	traversalString()

	// 替换字符串
	changeString()

	// 类型转换 T()
	var t int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	t = int(math.Sqrt(float64(3*3 + 4*4)))
	fmt.Println(t)

	//练习 统计汉字个数
	count := count("abcd你好世界")
	fmt.Println(count)
}

func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

//统计汉字个数
func count(str string) int {
	count := 0
	for _, r := range str {
		if len(string(r)) >= 3 {
			count++
		}
	}
	return count
}
