package main

import "testing"

// 类型别名
type MyInt = int64

func TestImplicit(t *testing.T) {

	var a int32 = 1
	var b int64

	// 显式类型转换
	b = int64(a)
	t.Log(a, b)

	var c MyInt
	c = b
	t.Log(c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1 不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	// string 数值类型
	var s string
	// 默认值为空字符串
	t.Log("*" + s + "*")
	t.Log(len(s))
}
