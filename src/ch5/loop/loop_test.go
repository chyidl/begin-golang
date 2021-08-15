package loop_test

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0

	// 循环
	for n < 5 {
		t.Log(n)
		n++
	}
}
