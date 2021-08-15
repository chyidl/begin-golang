package error_test

import (
	"errors"
	"testing"
)

// 区分错误类型
// 预先定义错误
var LessThanTwoError error = errors.New("n must be greater than 2")
var GreaterThanHunderedError error = errors.New("n must be less than 100")

func GetFibonacci(n int) ([]int, error) {
	// 快速失败
	if n < 2 {
		return nil, LessThanTwoError
	}

	if n > 100 {
		return nil, GreaterThanHunderedError
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

// 错误避免嵌套
func TestGetFibonacci(t *testing.T) {
	if list, err := GetFibonacci(10); err != nil {
		if err == LessThanTwoError {
			t.Error("Need a larger number")
		}

		if err == GreaterThanHunderedError {
			t.Error("Need a larger number")
		}
	} else {
		t.Log(list)
	}
}
