package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {

	var a int = 1
	var b int = 1

	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}

// Golang常量定义
// 快速设置连续值
func TestConstValue(t *testing.T) {
	const (
		Monday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	t.Log(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	const (
		Open = 1 << iota
		Close
		Pending
	)
	t.Log(Open, Close, Pending)
}
