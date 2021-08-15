package package_test

import (
	// src 默认属于gopath 路径一部分
	"ch15/series"
	"fmt"
	"testing"
)

func init() {
	fmt.Println("init func")
}

func init() {
	fmt.Println("init 2 func")
}

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(10))
	t.Log(series.Square(5))
}
