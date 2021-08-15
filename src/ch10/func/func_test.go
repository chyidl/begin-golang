package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 装饰器模式
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFunc(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

// 函数传递可变参数
func Sum(ops ...int) int {
	s := 0
	for _, op := range ops {
		s += op
	}
	return s
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(4, 5))
}

func Clear() {
	fmt.Println("Clear resources")
}

func TestDefer(t *testing.T) {
	// 匿名函数
	defer func() {
		t.Log("Clear resources")
	}()

	t.Log("Started")
	// 程序异常终端
	panic("Fatal error") // defer仍然会执行
}
