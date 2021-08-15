package if_test

import "testing"

func TestIf(t *testing.T) {
	// golang 支持方法多返回值
	if a := 1 == 1; a {
		t.Log("1==1")
	}
}
