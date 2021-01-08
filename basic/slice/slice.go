package main

import "fmt"

func main() {
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0, 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10

	//要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断
	var a1 []int             //len(a1)==0;cap(a1)=0;a1==nil
	a2 := []int{}            //len(a2)==0;cap(a2)=0;a2!=nil
	a3 := make([]int, 2, 10) //len(a3)==0;cap(a3)=0;a3!=nil
	fmt.Printf("len(a1)==%d;cap(a1)=%d;a1==%v", len(a1), cap(a1), a1)
	fmt.Printf("len(a2)==%d;cap(a2)=%d;a2==%v", len(a2), cap(a2), a2)
	fmt.Printf("len(a3)==%d;cap(a3)=%d;a3==%v", len(a3), cap(a3), a3)

	//切片的赋值拷贝
	s1 := make([]int, 3)
	s2 := s1 //将s1直接赋值给s2, s1与s2指向同一个数组
	s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)

	//切片遍历
	s := []int{1, 3, 5}
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
	for index, value := range s {
		fmt.Println(index, value)
	}

	//append()方法为切片添加元素
	var arr1 []int
	arr1 = append(arr1, 1)       //1
	arr1 = append(arr1, 2, 3, 4) //1 2 3 4
	arr2 := []int{5, 6, 7}
	arr1 = append(arr1, arr2...) //1 2 3 4 5 6 7
	fmt.Println(arr1)

	//通过var声明的零值切片可以在append()函数直接使用，无需初始化
	var arr3 = make([]int, 0)
	arr3 = append(arr3, arr1...)
	fmt.Println(arr3)

	//append()添加元素和切片扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	var citySlice []string
	// 追加一个元素
	citySlice = append(citySlice, "北京")
	// 追加多个元素
	citySlice = append(citySlice, "上海", "广州", "深圳")
	// 追加切片
	a5 := []string{"成都", "重庆"}
	citySlice = append(citySlice, a5...)
	fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]

	//使用copy()函数复制切片	直接通过赋值的形式，因为都指向同一数组导致相互影响
	var citySlice2 = make([]string, 6, 6)
	copy(citySlice2, citySlice)
	citySlice[5] = "长沙"
	fmt.Println(citySlice2) //[北京 上海 广州 深圳 成都 重庆]	不影响

	//从切片中删除元素
	a6 := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a6 = append(a6[:2], a6[3:]...)
	fmt.Println(a6) //[30 31 33 34 35 36 37]

	//练习
	var aa = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		aa = append(aa, fmt.Sprintf("%v", i))
	}
	fmt.Println(aa) //[     0 1 2 3 4 5 6 7 8 9] 前五个元素为""
}
