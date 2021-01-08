package split

import (
	"reflect"
	"testing"
)

//TestSplit 测试函数名必须以Test开头 必须接受一个*testing.T类型参数
func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")         //程序输出的结果
	want := []string{"a", "b", "c"}    //期望的结果
	if !reflect.DeepEqual(want, got) { //因为slice不能比较直接 接主反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) //测试失败输出错误提示
	}
}

//TestMoreSplit 单元测试
func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}
