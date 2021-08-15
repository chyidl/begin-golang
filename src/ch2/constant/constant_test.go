package constant

import (
	"testing"
)

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

	// 连续位常量赋值
	const (
		Open = 1 << iota
		Close
		Pending
	)
	t.Log(Open, Close, Pending)
}
