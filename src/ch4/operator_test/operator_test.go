package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 5}
	c := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	t.Log(b == c)
}

// 按照位清0
func TestBitClear(t *testing.T) {
	const (
		Executable = 1 << iota
		Writable
		Readable
	)

	a := 7
	a = a &^ Readable // 清除读权限
	a = a &^ Writable // 清除写权限
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
