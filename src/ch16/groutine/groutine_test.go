package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// 协程调度
		go func(i int) {
			fmt.Println(i)
		}(i) // 值传递
	}
	time.Sleep(time.Microsecond * 100)
}
